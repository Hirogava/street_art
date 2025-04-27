package api

import (
	"encoding/json"
	"log"
	"net/http"
	"street-art/db"
	"street-art/models"
	"street-art/services/cookies"
)

func Deposit(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	money := r.FormValue("money")
	userId := cookies.GetUserIdCookie(r)

	transactionType := models.TransactionTypeDeposit

	updatedBalance, err := manager.Deposit(userId, money, transactionType)
	if err != nil {
		log.Println("Ошибка при попытке пополнить баланс", err)
		http.Error(w, `{"message": "Ошибка пополнения баланса"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"new_balance": updatedBalance,
	})
	
}

func Purchase (w http.ResponseWriter, r *http.Request, manager *db.Manager) {
	userId := cookies.GetUserIdCookie(r)

	var order models.Orders
	order.UserId = userId

	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Println("Ошибка при декодировании данных заказа", err)
		http.Error(w, `{"message": "Ошибка при декодировании данных заказа"}`, http.StatusInternalServerError)
		return
	}

	transactionType := models.TransactionTypeWithdraw

	orderId, err := manager.Purchase(order, transactionType)
	if err != nil {
		if err.Error() == "not enough money" {
			http.Error(w, `{"message": "Недостаточно средств"}`, http.StatusPaymentRequired)
			return
		}
		log.Println("Ошибка при оформлении заказа", err)
		http.Error(w, `{"message": "Ошибка при оформлении заказа"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "ok",
		"order_id": orderId,
	})
}