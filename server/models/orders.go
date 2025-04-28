package models

type ProductOrder struct {
    ProductId   int               `json:"product_id"`
    ProductName string 			`json:"product_name"`
    Count       int               `json:"count"`
    Price       float64          `json:"price"`
}

type Orders struct {
    OrderId    int              `json:"order_id"`
    UserId     int              `json:"user_id"`
    Products   []ProductOrder   `json:"products"`
    TotalPrice float64          `json:"total_price"`
    OrderDate  string           `json:"order_date"`
    Status     string           `json:"status"`
}

type UserOrders struct {
    Orders []Orders
}