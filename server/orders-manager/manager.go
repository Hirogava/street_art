package ordersmanager

import (
	"context"
	"log"
	"street-art/db"
	"street-art/models"
	"sync"
	"time"
)

func OrderProcessing(ctx context.Context, orderId int, manager *db.Manager, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(time.Second * 10):
		err := manager.UpdateOrderStatus(orderId, models.OrderDone)
		if err != nil {
			log.Println("Ошибка обновления статуса заказа:", err)
		} else {
			log.Printf("Заказ %d успешно обновлен до статуса 'отправлен'\n", orderId)
		}
	case <-ctx.Done():
		log.Println("Контекст отменён для доставки заказа:", ctx.Err())
	}
}

func OrderDelivered(ctx context.Context, orderId int, manager *db.Manager, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(time.Second * 10):
		err := manager.UpdateOrderStatus(orderId, models.OrderDelivered)
		if err != nil {
			log.Println("Ошибка обновления статуса заказа:", err)
		} else {
			log.Printf("Заказ %d успешно обновлен до статуса 'доставлен'\n", orderId)
		}
	case <-ctx.Done():
		log.Println("Контекст отменён для доставки заказа:", ctx.Err())
	}
}