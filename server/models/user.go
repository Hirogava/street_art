package models

type User struct {
	Id int 				`json:"id"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
	Phone string 		`json:"phone"`
	Address string 		`json:"address"`
	Balance float64 	`json:"balance"`
}