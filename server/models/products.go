package models

type Product struct {
	Id int 				`json:"id"`
	Name string 		`json:"name"`
	Description string 	`json:"description"`
	Price float64 		`json:"price"`
	Category string 	`json:"category"`
	Stock int 			`json:"stock"`
	ImageUrl string 	`json:"image_url"`
	Brand string 		`json:"brand"`
}