package db

import (
	"database/sql"
	"fmt"
	"street-art/models"
)

func (manager *Manager) GetAllCart(userId int) (models.AllProductsFromCart, error) {
	var cart models.AllProductsFromCart

	query := `SELECT 
				c.id,
				p.id,
				p.name,
				p.description,
				p.price,
				p.stock,
				p.image_url,
				c.count
				FROM cart c
				JOIN products p ON c.product_id = p.id
				WHERE c.user_id = $1`
	rows, err := manager.Conn.Query(query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.AllProductsFromCart{}, fmt.Errorf("empty")
		}
		return models.AllProductsFromCart{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var productFromCart models.ProductFromCart
		err := rows.Scan(&productFromCart.CartId, &productFromCart.ProductId, &productFromCart.Name, &productFromCart.Description, &productFromCart.Price, &productFromCart.Stock, &productFromCart.ImageUrl, &productFromCart.Count)
		if err != nil {
			return models.AllProductsFromCart{}, err
		}
		productFromCart.TotalPrice = productFromCart.Price * float64(productFromCart.Count)
		cart.TotalPrice += productFromCart.TotalPrice
		cart.Count+=productFromCart.Count
		cart.ProductsFromCart = append(cart.ProductsFromCart, productFromCart)
	}

	cart.UserId = userId

	return cart, nil
}

func (manager *Manager) ClearCart(user_id int) error {
	query := `DELETE FROM cart WHERE user_id = $1`
	_, err := manager.Conn.Exec(query, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (manager *Manager) DeleteProductFromCart(cartId int, productId int) error {
	query := `DELETE FROM cart WHERE id = $1 AND product_id = $2`
	_, err := manager.Conn.Exec(query, cartId, productId)
	if err != nil {
		return err
	}

	return nil
}