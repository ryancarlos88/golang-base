package app

import (
	"database/sql"

	"github.com/ryancarlos88/golang-base/internal/infra/database"
	"github.com/ryancarlos88/golang-base/internal/infra/web/handler/customer"
	"github.com/ryancarlos88/golang-base/internal/infra/web/webserver"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/create"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/delete"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/update"
)

type App struct {
	ws *webserver.WebServer
	db *sql.DB
}

func NewApp(ws *webserver.WebServer, db *sql.DB) *App {
	return &App{ws, db}
}
func (a *App) Run() {
	repo := database.NewCustomerRepository(a.db)
	createUsecase := create.NewCreateCustomerUsecase(repo)
	updateUsecase := update.NewUpdateCustomerUsecase(repo)
	deleteUsecase := delete.NewDeleteCustomerUsecase(repo)
	handler := customer.NewCustomerHandler(createUsecase, updateUsecase, deleteUsecase)
	a.ws.AddHandler("/customers", "POST", handler.CreateCustomer)
	a.ws.AddHandler("/customers/{customerId}", "PUT", handler.UpdateCustomer)
	a.ws.AddHandler("/customers/{customerId}", "DELETE", handler.DeleteCustomer)
	a.ws.Start()
}
