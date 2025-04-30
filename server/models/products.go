package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	CategoryId  int
	Stock       int     `json:"stock"`
	ImageUrl    string  `json:"image_url"`
	Brand       string  `json:"brand"`
	BrandId     int
}

type ProductToOrder struct {
	OrderId     int     `json:"order_id"`
	ProductId   int     `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	TotalPrice  float64 `json:"total_price"`
	Count       int     `json:"stock"`
	ImageUrl    string  `json:"image_url"`
}

type ProductsToOrder struct {
	OrderId         int              `json:"order_id"`
	Status          string           `json:"status"`
	ProductsToOrder []ProductToOrder `json:"products_to_order"`
}

type ProductToCart struct {
	Id     int `json:"id"`
	Count  int `json:"count"`
	UserId int `json:"user_id"`
}

type ProductFromCart struct {
	CartId      int     `json:"cart_id"`
	ProductId   int     `json:"product_id"`
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
	Stock       int     `json:"stock"`
	TotalPrice  float64 `json:"total_price"`
}

type AllProductsFromCart struct {
	UserId           int               `json:"user_id"`
	ProductsFromCart []ProductFromCart `json:"products_from_cart"`
	TotalPrice       float64           `json:"total_price"`
	Count            int               `json:"count"`
}