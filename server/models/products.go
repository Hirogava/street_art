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

type ProductToCart struct {
	Id int 				`json:"id"`
	Count int 			`json:"count"`
	UserId int 			`json:"user_id"`
}