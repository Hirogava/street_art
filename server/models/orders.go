package models

type ProductOrder struct {
    ProductId int               `json:"product_id"`
    Count     int               `json:"count"`
}

type Orders struct {
    UserId    int               `json:"user_id"`
    Products   []ProductOrder   `json:"products"`
    TotalPrice float64          `json:"total_price"`
}