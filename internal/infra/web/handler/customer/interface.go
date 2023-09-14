package customer

import (
	"context"

	cd "github.com/ryancarlos88/golang-base/internal/usecase/customer/create/dto"
	ld "github.com/ryancarlos88/golang-base/internal/usecase/customer/list/dto"
	rd "github.com/ryancarlos88/golang-base/internal/usecase/customer/read/dto"
	ud "github.com/ryancarlos88/golang-base/internal/usecase/customer/update/dto"
)

type CreateCustomerUsecase interface {
	Execute(context.Context, cd.CreateCustomerInput) (*cd.CreateCustomerOutput, error)
}

type UpdateCustomerUsecase interface {
	Execute(context.Context, ud.UpdateCustomerInput) (*ud.UpdateCustomerOutput, error)
}

type DeleteCustomerUsecase interface {
	Execute(context.Context, string) error
}

type ReadCustomerUsecase interface {
	Execute(context.Context, string) (*rd.ReadCustomerOutput, error)
}

type ListCustomerUsecase interface {
	Execute(context.Context, ld.ListCustomersInput) ([]ld.ListCustomersOutput, error)
}
