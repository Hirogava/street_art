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