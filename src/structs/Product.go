package structs

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code    string
	Product string
	Price   uint
	Unit    string
}
