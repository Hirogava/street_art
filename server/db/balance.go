package db

import (
	"strconv"
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