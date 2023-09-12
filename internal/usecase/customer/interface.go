package customer

import (
	"context"

	"github.com/ryancarlos88/golang-base/internal/entity"
)

type CreateCustomerUsecaseRepository interface {
	Create(ctx context.Context, customer entity.Customer) error
}

type DeleteCustomerUsecaseRepository interface {
	Delete(ctx context.Context, customerID string) error
}

type UpdateCustomerUsecaseRepository interface {
	Update(ctx context.Context, customer entity.Customer) error
}