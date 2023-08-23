package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gertd/go-pluralize"
	driver2 "github.com/xiaoshouchen/go-model-gen/driver"
	"github.com/xiaoshouchen/go-model-gen/pkg"
	"github.com/xiaoshouchen/go-model-gen/vars"
	"go/format"
	"html/template"
	"log"
	"os"
	"syscall"
)

var dataSource = flag.String("f", "data_source.json", "文件地址")
var outputSource = flag.String("o", "./model/", "文件输出位置")

//go:embed tpl/model.tpl
var model string

//go:embed tpl/field.tpl
var field string

//go:embed tpl/insert.tpl
var insert string

func main() {
	var configs = make([]vars.DatabaseConfig, 0)
	getConfig(&configs)
	plur := pluralize.NewClient()
	for _, config := range configs {
		inst := driver2.GetInstance(config)
		schemas := inst.GetTableStructure(config.Scheme, config.Tables)
		err := os.Mkdir(*outputSource, 0777)
		if err != nil && !os.IsExist(err) {
			fmt.Println(err)
			return
		}
		if len(schemas) > 0 {
			for _, s := range schemas {
				t, err := template.New("modelTpl").Funcs(template.FuncMap{
					"lowCamel":  pkg.LineToLowCamel,
					"upCamel":   pkg.LineToUpCamel,
					"singular":  plur.Singular,
					"plural":    plur.Plural,
					"transType": inst.TransType,
				}).Parse(insert + field + model)
				if err != nil {
					log.Fatal(err)
					return
				}
				var buf = new(bytes.Buffer)
				if err != nil {
					log.Fatal(err)
					return
				}
				err = t.ExecuteTemplate(buf, "modelTpl", s)
				source, err := format.Source(buf.Bytes())
				if err != nil {
					log.Fatal(err)
					return
				}
				fileName := *outputSource + plur.Singular(s.TableName) + "_gen.go"
				err = os.WriteFile(fileName, source, 0664)
				if err != nil {
					fmt.Println(err)
					return
				}
				path := *outputSource + plur.Singular(s.TableName) + ".go"
				if exists, err := pkg.FileExists(path); !exists || err != nil {
					err = os.WriteFile(path, []byte("package model"), 0664)
					if err != nil {
						return
					}
				}
			}
		}
		syscall.Exec("gofmt", []string{"-w", *outputSource}, []string{})
	}

}

func getConfig(j *[]vars.DatabaseConfig) {
	flag.Parse()
	tables, err := os.ReadFile(*dataSource)
	if err != nil {
		fmt.Println("please create a data_source.json first")
	}
	var configMap = make(map[string][]vars.DatabaseConfig)
	err = json.Unmarshal(tables, &configMap)
	if configs, ok := configMap["databases"]; ok {
		*j = configs
	}
}
