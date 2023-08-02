package formatter

import "gorm.io/gorm"

type Name struct {
	Test int `json:"test"`
}

type NameModel struct {
	db *gorm.DB
}

func NewNameModel(db *gorm.DB) *NameModel {
	return &NameModel{db: db}
}

func (m *NameModel) Insert(name Name) {
	m.db.Create(name)
}

func (m *NameModel) MultiInsert(nameSlice []Name) {

}

func (m *NameModel) Update() {

}
