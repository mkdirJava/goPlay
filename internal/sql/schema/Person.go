package schema

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name  string
	Age   int16
	Time  string
	Likes int16
}
