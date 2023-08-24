# go-model-gen

> a tool generating golang database model struct matching all fields    
> warning: this project depends on gorm.io

## example
you will get some code like below

```go
// DON'T EDIT !

package model

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Admin struct {
	Id             int64        `json:"id" gorm:"primaryKey"` //
	Username       string       `json:"username" `            //用户名
	Name           string       `json:"name" `                //用户姓名
	Password       string       `json:"password" `            //
	PasswordModify int64        `json:"password_modify" `     //
	DeletedAt      sql.NullTime `json:"deleted_at" `          //
	CreatedAt      sql.NullTime `json:"created_at" `          //
	UpdatedAt      sql.NullTime `json:"updated_at" `          //
}

func (Admin) TableName() string {
	return "admins"
}

type AdminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{db: db}
}

// Insert 插入，支持批量插入，单个插入，支持分批插入，支持冲突更新或者报错
// allowUpdate 允许报错更新
// batchSize 如果数据过多，比如几万个，则需要拆分多个，等于小于0则不分批
func (r *AdminRepo) Insert(allowUpdate bool, batchSize int, insertSlice ...*Admin) (int64, error) {
	db := r.db
	if allowUpdate {
		db = db.Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
					{Name: "id"},
				},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"username":        db.Raw("values(username)"),
					"name":            db.Raw("values(name)"),
					"password":        db.Raw("values(password)"),
					"password_modify": db.Raw("values(password_modify)"),
					"deleted_at":      db.Raw("values(deleted_at)"),
					"created_at":      db.Raw("values(created_at)"),
					"updated_at":      db.Raw("values(updated_at)"),
				}),
			},
		)
	}
	if batchSize > 0 {
		db = db.CreateInBatches(insertSlice, batchSize)

	} else {
		db = db.Create(insertSlice)
	}
	return db.RowsAffected, db.Error
}

func (r *AdminRepo) getAllFields() []string {
	return []string{
		"id",
		"username",
		"name",
		"password",
		"password_modify",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

// omit 过滤自己不想要的字段
func (r *AdminRepo) omit(filter []string) []string {
	fields := r.getAllFields()
	result := make([]string, 0, len(fields))
	filterSet := make(map[string]bool)

	// 将需要过滤的值添加到 filterSet 中
	for _, v := range filter {
		filterSet[v] = true
	}

	// 遍历原始切片，将不在 filterSet 中的值添加到结果切片中
	for _, v := range fields {
		if _, ok := filterSet[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

```

## install

```shell
go install github.com/xiaoshouchen/go-model-gen@latest
```

## how to use

create a json file,and a string array

// data_source.json

```json
[
  "table1",
  "table2"
]

```