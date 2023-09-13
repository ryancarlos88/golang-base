package list

import (
	"context"
	"errors"
	"fmt"

	"github.com/ryancarlos88/golang-base/internal/usecase/customer"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/list/dto"
)

type ListCustomersUsecase struct {
	repo customer.ListCustomerUsecaseRepository
}

func NewListCustomerUsecase(repo customer.ListCustomerUsecaseRepository) *ListCustomersUsecase {
	return &ListCustomersUsecase{repo}
}

var ErrListingCustomers = errors.New("error listing customers")

func (u *ListCustomersUsecase) Execute(ctx context.Context, input dto.ListCustomersInput) ([]dto.ListCustomersOutput, error) {
	res, err := u.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%w]: %v", ErrListingCustomers, err)
	}

	output := make([]dto.ListCustomersOutput, len(res))
	for i := 0; i < len(output); i++ {
		output[i].ID = res[i].ID
		output[i].Name = res[i].Name
		output[i].Phone = res[i].Phone
		output[i].Email = res[i].Email
	}
	return output, nil
}
