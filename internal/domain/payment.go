package domain

import "time"

type Payment struct {
	Id      *string
	OrderId *string
	Method  *string
	Status  *string
	PaidAt  *time.Time
}

func (Payment) TableName() string {
	return "payment"
}
