package admin

import "github.com/jinzhu/gorm"

type AuthGroup struct {
	gorm.Model
	Title  string
	Status int
	Rules  string
}
