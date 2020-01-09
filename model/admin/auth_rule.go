package admin

import "github.com/jinzhu/gorm"

type AuthRule struct {
	gorm.Model
	Name      string
	Title     string
	Type      int
	Status    int
	Css       string
	Condition string
	Pid       int
	Sort      int
}
