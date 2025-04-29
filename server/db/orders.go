package db

import (
	"database/sql"
	"fmt"
	"street-art/models"
	"time"
)

func (manager *Manager) GetAllOrders(userId int) (models.UserOrders, error) {
	var userOrders models.UserOrders

	rows, err := manager.Conn.Query("SELECT id, order_date, status, total_price FROM orders WHERE user_id = $1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserOrders{}, fmt.Errorf("no orders")
		}
		return models.UserOrders{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Orders

		order.UserId = userId
		err := rows.Scan(&order.OrderId, &order.OrderDate, &order.Status, &order.TotalPrice)
		if err != nil {
			return models.UserOrders{}, err
		}

		parsedTime, err := time.Parse(time.RFC3339Nano, order.OrderDate)
		if err != nil {
			return models.UserOrders{}, err
		}
		order.OrderDate = parsedTime.Format("02.01.2006 15:04")

		err = manager.GetProductsForOrders(&order.Products, order.OrderId)
		if err != nil {
			return models.UserOrders{}, err
		}

		userOrders.Orders = append(userOrders.Orders, order)
	}

	return userOrders, nil
}

func (manager *Manager) GetProductsForOrders(productOrders *[]models.ProductOrder, orderId int) error {
	rows, err := manager.Conn.Query(`SELECT
										p.name,
										o.product_id,
										o.quantity,
										o.price
										FROM order_items AS o
										JOIN products AS p ON p.id = o.product_id
										WHERE order_id = $1`, orderId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var productOrder models.ProductOrder
		err := rows.Scan(&productOrder.ProductName, &productOrder.ProductId, &productOrder.Count, &productOrder.Price)
		if err != nil {
			return err
		}
		*productOrders = append(*productOrders, productOrder)
	}

	return nil
}

func (manager *Manager) GetOrderDetails(orderId int) (models.ProductsToOrder, error) {
	var products models.ProductsToOrder

	var status string
	err := manager.Conn.QueryRow("SELECT status FROM orders WHERE id = $1", orderId).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ProductsToOrder{}, fmt.Errorf("заказ не найден")
		}
		return models.ProductsToOrder{}, err
	}
	products.Status = status

	rows, err := manager.Conn.Query(`SELECT
										o.product_id,
										p.name,
										p.description,
										o.price,
										o.quantity,
										p.image_url
										FROM order_items AS o
										JOIN products AS p ON p.id = o.product_id
										WHERE order_id = $1`, orderId)
	if err != nil {
		return models.ProductsToOrder{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.ProductToOrder
		err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.TotalPrice, &product.Count, &product.ImageUrl)
		if err != nil {
			return models.ProductsToOrder{}, err
		}
		products.ProductsToOrder = append(products.ProductsToOrder, product)
	}
	products.OrderId = orderId

	return products, nil
}

func (manager *Manager) UpdateOrderStatus(orderId int, status string) error {
	_, err := manager.Conn.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, orderId)
	return err
}

func (manager *Manager) GetAllOrdersForAdmin() (models.UserOrders, error) {
	var userOrders models.UserOrders

	rows, err := manager.Conn.Query("SELECT id, order_date, status, total_price, user_id FROM orders")
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserOrders{}, fmt.Errorf("no orders")
		}
		return models.UserOrders{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Orders

		err := rows.Scan(&order.OrderId, &order.OrderDate, &order.Status, &order.TotalPrice, &order.UserId)
		if err != nil {
			return models.UserOrders{}, err
		}

		parsedTime, err := time.Parse(time.RFC3339Nano, order.OrderDate)
		if err != nil {
			return models.UserOrders{}, err
		}
		order.OrderDate = parsedTime.Format("02.01.2006 15:04")

		err = manager.GetProductsForOrders(&order.Products, order.OrderId)
		if err != nil {
			return models.UserOrders{}, err
		}

		userOrders.Orders = append(userOrders.Orders, order)
	}

	return userOrders, nil
}