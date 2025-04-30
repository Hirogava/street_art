package models

type Category struct {
	Id   int  				`json:"id"`
	Name string 			`json:"name"`
	Description string 		`json:"description"`
	ImageUrl string 		`json:"image_url"`
}

type CategoryMini struct {
	Id   int  				`json:"id"`
	Name string 			`json:"name"`
}