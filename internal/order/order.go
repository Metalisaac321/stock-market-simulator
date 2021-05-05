package order

import (
	"context"
)

// Order
type Order struct {
	Id                 string
	OrderTypeOperation string
	IssuerName         string
	TotalShares        uint
	SharePrice         uint
	OrderDate          string
	AccountId          string
}

// Repository
type OrderRepository interface {
	Save(ctx context.Context, order Order) error
}
