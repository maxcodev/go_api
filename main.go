package main

import (
	s "github.com/maxcodev/go_api/src/structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:12m08s95c@@@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&s.Product{})

	// data := &s.Product{}

	// data.Code = "D54"
	// data.Product = "Pinza multiuso"
	// data.Price = 800
	// data.Unit = "Pieza"
	//Create
	//db.Create(data)

	/*
	* Instanciamos el objeto product y con db.First realizamos la consulta
	* en la tabla Product que recibe dos parametro, el modelo y el id opcional
	 */
	var product s.Product
	db.First(&product, 2)

	product.Code = "D56"
	product.Price = 999
	product.Unit = "Pieza"

	/*
	* Update
	 */
	db.Model(&product).Updates(map[string]interface{}{
		"code":  product.Code,
		"price": product.Price,
		"unit":  product.Unit,
	})
}
