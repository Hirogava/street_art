package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"street-art/models"
)

func (manager *Manager) Deposit(userId int, money string, transactionType string) (float64, error) {
	floatMoney, err := strconv.ParseFloat(money, 64)
	if err != nil {
		return 0, err
	}
	
	var updatedBalance float64

	err = manager.Conn.QueryRow("UPDATE users SET balance = balance + $1 WHERE id = $2 RETURNING balance", floatMoney, userId).Scan(&updatedBalance)
	if err != nil {
		return 0, err
	}

	_, err = manager.Conn.Exec("INSERT INTO balance_transactions (user_id, transaction_type, amount) VALUES ($1, $2, $3)", userId, transactionType, floatMoney)
	if err != nil {
		return 0, err
	}

	return updatedBalance, nil
}

func (manager *Manager) GetUserBalance(userId int) (float64, error) {
	var balance float64

	err := manager.Conn.QueryRow("SELECT balance FROM users WHERE id = $1", userId).Scan(&balance)
	if err != nil {
		return 0, err
	}
	
	return balance, nil
}

func (manager *Manager) Purchase(order models.Orders, transactionType string) (int, error) {
    tx, err := manager.Conn.Begin()
    if err != nil {
        return 0, err
    }
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    var balance float64
    err = tx.QueryRow("SELECT balance FROM users WHERE id = $1", order.UserId).Scan(&balance)
    if err != nil {
        return 0, err
    }

    if order.TotalPrice > balance {
        return 0, fmt.Errorf("not enough money")
    }

    _, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", order.TotalPrice, order.UserId)
    if err != nil {
        return 0, err
    }

    _, err = tx.Exec("INSERT INTO balance_transactions (user_id, transaction_type, amount) VALUES ($1, $2, $3)", order.UserId, transactionType, order.TotalPrice)
    if err != nil {
        return 0, err
    }

    var orderId int
    err = tx.QueryRow("INSERT INTO orders (user_id, total_price) VALUES ($1, $2) RETURNING id", order.UserId, order.TotalPrice).Scan(&orderId)
    if err != nil {
        return 0, err
    }

    err = manager.CreateOrderItemsTx(tx, orderId, order.Products)
    if err != nil {
        return 0, err
    }

    _, err = tx.Exec("DELETE FROM cart WHERE user_id = $1", order.UserId)
    if err != nil {
        return 0, err
    }

    err = tx.Commit()
    if err != nil {
        return 0, err
    }

    return orderId, nil
}

func (manager *Manager) CreateOrderItemsTx(tx *sql.Tx, orderId int, orders []models.ProductOrder) error {
	for _, productOrder := range orders {
		price, err := manager.GetItemPrice(productOrder.ProductId, productOrder.Count)
		if err != nil {
			return err
		}

		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderId, productOrder.ProductId, productOrder.Count, price)
		if err != nil {
			return err
		}

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1", productOrder.Count, productOrder.ProductId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (manager *Manager) GetItemPrice(productId int, count int) (float64, error) {
	var price float64

	err := manager.Conn.QueryRow("SELECT price FROM products WHERE id = $1", productId).Scan(&price)
	if err != nil {
		return 0, err
	}

	return price * float64(count), nil
}