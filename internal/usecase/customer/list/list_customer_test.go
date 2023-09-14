package list

import (
	"context"
	"errors"
	"testing"

	"github.com/ryancarlos88/golang-base/internal/entity"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/list/dto"
	"github.com/stretchr/testify/assert"
)

type customerRepoMock struct {
	returnsList []entity.Customer
	returnsErr  error
}

func (m *customerRepoMock) List(_ context.Context, _, _ int) ([]entity.Customer, error) {
	return m.returnsList, m.returnsErr
}

func TestExecute(t *testing.T) {
	id := "some_random_id"
	name := "teste"
	phone := "1122223333"
	email := "test"
	returnedList := []entity.Customer{
		{ID: id, Name: name, Phone: phone, Email: email},
	}

	expected := []dto.ListCustomersOutput{
		{ID: id, Name: name, Phone: phone, Email: email},
	}
	repo := &customerRepoMock{returnsList: returnedList}
	uc := NewListCustomerUsecase(repo)

	res, err := uc.Execute(context.Background(), dto.ListCustomersInput{})
	assert.Nil(t, err)
	assert.Equal(t, res, expected)

}

func TestExecute_Error(t *testing.T) {

	repo := &customerRepoMock{returnsErr: errors.New("some random err")}
	uc := NewListCustomerUsecase(repo)

	_, err := uc.Execute(context.Background(), dto.ListCustomersInput{})
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrListingCustomers)

}
