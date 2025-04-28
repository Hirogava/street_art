package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"street-art/db"
	"street-art/models"
	"street-art/services/cookies"
	"time"
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

func Purchase(w http.ResponseWriter, r *http.Request, manager *db.Manager) {
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

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// wg := &sync.WaitGroup{}

	// wg.Add(2)
	// go ordersmanager.OrderProcessing(ctx, orderId, manager, wg)
	// go ordersmanager.OrderDelivered(ctx, orderId, manager, wg)
	// wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    
    go func(ctx context.Context, orderId int, manager *db.Manager) {
        defer cancel()
        
        select {
        case <-time.After(time.Second * 10):
            err := manager.UpdateOrderStatus(orderId, models.OrderDone)
            if err != nil {
                log.Println("Ошибка обновления статуса заказа:", err)
            } else {
                log.Printf("Заказ %d успешно обновлен до статуса 'отправлен'\n", orderId)
            }
        case <-ctx.Done():
            log.Println("Обновление статуса заказа отменено:", ctx.Err())
        }
    }(ctx, orderId, manager)
}