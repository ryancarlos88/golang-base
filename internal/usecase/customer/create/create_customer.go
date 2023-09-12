package create

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryancarlos88/golang-base/internal/entity"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/create/dto"
)

type CreateCustomerUsecase struct {
	repo customer.CreateCustomerUsecaseRepository
}

func NewCreateCustomerUsecase(repo customer.CreateCustomerUsecaseRepository) *CreateCustomerUsecase {
	return &CreateCustomerUsecase{repo: repo}
}

var ErrCreatingCustomer = errors.New("error creating customer")

func (u *CreateCustomerUsecase) Execute(ctx context.Context, input dto.CreateCustomerInput) (*dto.CreateCustomerOutput, error) {
	customer := entity.NewCustomer(input.Name, input.Phone, input.Email)
	err := customer.Validate()
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("[%w]: %v", ErrCreatingCustomer, err)
	}

	out := &dto.CreateCustomerOutput{
		ID:    customer.ID,
		Name:  customer.Name,
		Phone: customer.Phone,
		Email: customer.Email,
	}

	return out, nil
}
