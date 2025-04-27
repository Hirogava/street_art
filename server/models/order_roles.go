package models

const (
	OrderWait = "в обработке"
	OrderDone = "отправлен"
	OrderDelivered = "доставлен"

	PaymentCard = "карта"
	PaymentSBP = "сбп"

	PaymentStatusWait = "ожидает"
	PaymentStatusPayed = "оплачен"

	TransactionTypeDeposit = "пополнение"
	TransactionTypeWithdraw = "списание"
)