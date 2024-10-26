package models

import "gorm.io/gorm"

// item struct
type Item struct {
	gorm.Model
	Title   string `json:"title"`
	Checked bool   `json:"checked"`
}
