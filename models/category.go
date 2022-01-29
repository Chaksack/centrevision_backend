package models

import "gorm.io/gorm"

type Category struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func (category *Category) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Category{}).Count(&total)
	return total
}

func (category *Category) Take(db *gorm.DB, limit int, offset int) interface{} {
	var categorys []Category
	db.Offset(offset).Limit(limit).Find(&categorys)
	return categorys
}
