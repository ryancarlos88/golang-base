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

func (r *CustomerRepository) Read(ctx context.Context, customerID string) (*entity.Customer, error) {
	return &entity.Customer{
		ID: customerID,
		Name: "to be implemented",
		Phone: "to be implemented",
		Email: "to be implemented",
	}, nil
}

func (r *CustomerRepository) List(ctx context.Context, pageNumber, pageSize int) ([]entity.Customer, error) {
	return []entity.Customer{
		{
			ID: "123", 
			Name: "to be implemented", 
			Phone: "to be implemented", 
			Email: "to be implemented",
		},
		{
			ID: "456", 
			Name: "to be implemented", 
			Phone: "to be implemented", 
			Email: "to be implemented",
		},
	}, nil
}
