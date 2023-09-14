package customer

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	createdto "github.com/ryancarlos88/golang-base/internal/usecase/customer/create/dto"
	"github.com/ryancarlos88/golang-base/internal/usecase/customer/list/dto"
	updatedto "github.com/ryancarlos88/golang-base/internal/usecase/customer/update/dto"
)

type CustomerHandler struct {
	createUsecase CreateCustomerUsecase
	updateUsecase UpdateCustomerUsecase
	deleteUsecase DeleteCustomerUsecase
	readUsecase   ReadCustomerUsecase
	listUsecase   ListCustomerUsecase
}

func NewCustomerHandler(create CreateCustomerUsecase,
	update UpdateCustomerUsecase,
	delete DeleteCustomerUsecase,
	read ReadCustomerUsecase,
	list ListCustomerUsecase) *CustomerHandler {
	return &CustomerHandler{create, update, delete, read, list}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var dto createdto.CreateCustomerInput

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.createUsecase.Execute(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var dto updatedto.UpdateCustomerInput

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dto.ID = chi.URLParam(r, "customerId")

	out, err := h.updateUsecase.Execute(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customerID = chi.URLParam(r, "customerId")

	err := h.deleteUsecase.Execute(r.Context(), customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}

func (h *CustomerHandler) ListCustomers(w http.ResponseWriter, r *http.Request) {
	var input dto.ListCustomersInput
	ps, err := strconv.Atoi(chi.URLParam(r, "page_size"))
	if err != nil {
		ps = 1
	}
	pn, err := strconv.Atoi(chi.URLParam(r, "page_number"))
	if err != nil {
		ps = 1
	}
	input.PageNumber = pn
	input.PageSize = ps

	out, err := h.listUsecase.Execute(context.Background(), input)

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CustomerHandler) ReadCustomer(w http.ResponseWriter, r *http.Request) {
	var customerID = chi.URLParam(r, "customerId")

	out, err := h.readUsecase.Execute(context.Background(), customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
