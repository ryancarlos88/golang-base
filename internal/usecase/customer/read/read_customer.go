package read

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryancarlos88/golang-base/internal/usecase/customer"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/read/dto"
)

type ReadCustomerUsecase struct {
	repo customer.ReadCustomerUsecaseRepository
}

func NewReadCustomerUsecase(repo customer.ReadCustomerUsecaseRepository) *ReadCustomerUsecase {
	return &ReadCustomerUsecase{repo}
}

var ErrReadingCustomer = errors.New("error reading customer")

func (u *ReadCustomerUsecase) Execute(ctx context.Context, customerID string) (*dto.ReadCustomerOutput, error) {
	customer, err := u.repo.Read(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("[%w] %v", ErrReadingCustomer, err)
	}

	out := &dto.ReadCustomerOutput{
		ID:    customer.ID,
		Name:  customer.Name,
		Phone: customer.Phone,
		Email: customer.Email,
	}
	return out, nil
}
