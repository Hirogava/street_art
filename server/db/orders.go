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