package models

import "gorm.io/gorm"

type Statement struct {
	gorm.Model
	Keywords  []Keyword
	Variables []Variable
}

type Keyword struct {
	Var string
}

type Variable struct {
	Var   string
	Type  string
	Value string
}
