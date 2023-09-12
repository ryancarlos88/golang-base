package entity

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Customer struct {
	ID    string
	Name  string
	Phone string
	Email string
}

func NewCustomer(name, phone, email string) Customer {
	id := uuid.New().String()
	customer := Customer{
		ID: id, Name: name, Phone: phone, Email: email,
	}
	return customer
}

var ErrInvalidPhone = errors.New("phone number is invalid")

func (c *Customer) Validate() error {
	if len(c.Phone) < 10 {
		return fmt.Errorf("%w", ErrInvalidPhone)
	}
	return nil
}
