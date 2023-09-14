package read

import (
	"context"
	"errors"
	"testing"

	"github.com/ryancarlos88/golang-base/internal/entity"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/read/dto"
	"github.com/stretchr/testify/assert"
)

type customerRepoMock struct {
	customer *entity.Customer
	err      error
}

func (m *customerRepoMock) Read(_ context.Context, _ string) (*entity.Customer, error) {
	return m.customer, m.err
}

func TestExecute(t *testing.T) {
	id := "id"
	name := "test"
	phone := "1122223333"
	email := "test@test.com"

	expected := &dto.ReadCustomerOutput{ID: id, Name: name, Phone: phone, Email: email}

	m := &customerRepoMock{customer: &entity.Customer{ID: id, Name: name, Phone: phone, Email: email}}

	uc := NewReadCustomerUsecase(m)

	out, err := uc.Execute(context.Background(), "1")
	assert.Nil(t, err)
	assert.Equal(t, expected, out)
}

func TestExecute_Error(t *testing.T) {
	m := &customerRepoMock{err: errors.New("some err")}

	uc := NewReadCustomerUsecase(m)

	_, err := uc.Execute(context.Background(), "1")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrReadingCustomer)
}