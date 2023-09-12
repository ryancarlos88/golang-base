package database

import (
	"context"
	"database/sql"

	"github.com/ryancarlos88/golang-base/internal/entity"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) Create(ctx context.Context, customer entity.Customer) error {
	return nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer entity.Customer) error {
	return nil
}

func (r *CustomerRepository) Delete(ctx context.Context, customerID string) error {
	return nil
}
