package create

import (
	"context"
	"errors"
	"testing"

	"github.com/ryancarlos88/golang-base/internal/entity"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/create/dto"
	"github.com/stretchr/testify/assert"
)

type customerRepoMock struct {
	returnedErr error
}

func (m *customerRepoMock) Create(_ context.Context, _ entity.Customer) error {
	return m.returnedErr
}

func TestExecute(t *testing.T) {
	name := "test"
	phone := "1122223333"
	email := "test@test.com"

	expected := &dto.CreateCustomerOutput{Name: name, Phone: phone, Email: email}

	repo := &customerRepoMock{}

	uc := NewCreateCustomerUsecase(repo)

	res, _ := uc.Execute(context.Background(), dto.CreateCustomerInput{Name: name, Phone: phone, Email: email})

	assert.Equal(t, expected, res)
}

func TestExecute_Error(t *testing.T) {
	cases := []struct {
		testName                 string
		name, phone, email       string
		testExpectedErr, repoErr error
	}{
		{
			testName:        "should return invalid phone err",
			name:            "teste",
			phone:           "11222333",
			email:           "test@test.com",
			testExpectedErr: entity.ErrInvalidPhone,
		},
		{
			testName:        "should return invalid phone err",
			name:            "teste",
			phone:           "1122223333",
			email:           "teste@test.com",
			repoErr:         errors.New("some error"),
			testExpectedErr: ErrCreatingCustomer,
		},
	}
	for _, tc := range cases {
		repo := &customerRepoMock{returnedErr: tc.repoErr}

		uc := NewCreateCustomerUsecase(repo)

		_, err := uc.Execute(context.Background(), dto.CreateCustomerInput{Name: tc.name, Phone: tc.phone, Email: tc.email})

		assert.NotNil(t, err)
		assert.ErrorIs(t, err, tc.testExpectedErr)
	}
}
