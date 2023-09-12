package delete

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryancarlos88/golang-base/internal/usecase/customer"
)

type DeleteCustomerUsecase struct {
	repo customer.DeleteCustomerUsecaseRepository
}

func NewDeleteCustomerUsecase(repo customer.DeleteCustomerUsecaseRepository) *DeleteCustomerUsecase {
	return &DeleteCustomerUsecase{repo}
}

var ErrDeletingCustomer = errors.New("error deleting customer")
func (u *DeleteCustomerUsecase)Execute(ctx context.Context, customerID string) error {
	err := u.repo.Delete(ctx, customerID)
	if err != nil {
		return fmt.Errorf("[%w]: %v", ErrDeletingCustomer, err)
	}

	return nil
}