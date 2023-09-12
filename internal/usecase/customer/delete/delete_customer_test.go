package delete

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type deleteCustomerRepoMock struct {
	returnsErr error
}

func (m *deleteCustomerRepoMock) Delete(_ context.Context, _ string) error{
	return m.returnsErr
}

func TestExecute(t *testing.T){
	repo := &deleteCustomerRepoMock{}
	uc := NewDeleteCustomerUsecase(repo)

	err := uc.Execute(context.Background(), "customerId")

	assert.Nil(t, err)
}
func TestExecute_Error(t *testing.T){
	repo := &deleteCustomerRepoMock{returnsErr: errors.New("some error")}
	uc := NewDeleteCustomerUsecase(repo)

	err := uc.Execute(context.Background(), "customerId")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrDeletingCustomer)
}