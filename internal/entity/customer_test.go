package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	name := "test"
	phone := "1122223333"
	email := "test@test.com"
	c := NewCustomer(name, phone, email)

	assert.Equal(t, name, c.Name)
	assert.Equal(t, phone, c.Phone)
	assert.Equal(t, email, c.Email)
}

func TestValidate(t *testing.T) {
	name := "test"
	phone := "1122223333"
	email := "test@test.com"
	c := NewCustomer(name, phone, email)
	err := c.Validate()
	assert.Nil(t, err)
	assert.Equal(t, name, c.Name)
	assert.Equal(t, phone, c.Phone)
	assert.Equal(t, email, c.Email)
}

func TestValidate_Error(t *testing.T) {
	name := "test"
	phone := "112223333"
	email := "test@test.com"
	c := NewCustomer(name, phone, email)
	err := c.Validate()
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidPhone)
}
