package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"go/format"
	"os"

	"github.com/xiaoshouchen/go-model-gen/driver"
	"github.com/xiaoshouchen/go-model-gen/pkg"
)

func main() {
	tables, err := os.ReadFile("data_source.json")
	if err != nil {
		fmt.Println("please create a data_source.json first")
	}
	var j TableJson
	err = json.Unmarshal(tables, &j)
	if err != nil {
		return
	}
	schemas := driver.InitPostgres(j.Tables)

	err = os.Mkdir("model", 0777)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
		return
	}
	if len(schemas) > 0 {
		for _, s := range schemas {
			fileContent := fmt.Sprintf("//")
			fileContent = fmt.Sprintf("package model \n")
			fileContent += fmt.Sprintf("type %s struct { \n", pkg.LineToCamel(s.TableName))
			for _, f := range s.Fields {
				dataType := pkg.TransType(f.DataType)
				fileContent += fmt.Sprintf("%s %s  `db:\"%s\" json:\"%s\"` \n", pkg.LineToCamel(f.ColumnName), dataType, f.ColumnName, f.ColumnName)
			}
			fileContent += fmt.Sprintf("} \n")
			bytes := []byte(fileContent)
			bytes, err = format.Source(bytes)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = os.WriteFile("model/"+s.TableName+"_gen.go", bytes, 0664)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

type TableJson struct {
	Tables []string `json:"tables"`
}
