package update

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryancarlos88/golang-base/internal/entity"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/update/dto"
)

type UpdateCustomerUsecase struct {
	repo customer.UpdateCustomerUsecaseRepository
}

func NewUpdateCustomerUsecase(repo customer.UpdateCustomerUsecaseRepository) *UpdateCustomerUsecase {
	return &UpdateCustomerUsecase{repo: repo}
}

var ErrUpdatingCustomer = errors.New("error updating customer")

func (u *UpdateCustomerUsecase) Execute(ctx context.Context, input dto.UpdateCustomerInput) (*dto.UpdateCustomerOutput, error) {
	customer := entity.Customer{
		ID:    input.ID,
		Name:  input.Name,
		Phone: input.Phone,
		Email: input.Email,
	}
	err := customer.Validate()
	if err != nil {
		return nil, err
	}
	err = u.repo.Update(ctx, customer)
	if err != nil {
		return nil, fmt.Errorf("[%w]: %v", ErrUpdatingCustomer, err)
	}

	out := &dto.UpdateCustomerOutput{
		ID:    customer.ID,
		Name:  customer.Name,
		Phone: customer.Phone,
		Email: customer.Email,
	}

	return out, nil
}
