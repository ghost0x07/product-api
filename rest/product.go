package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"product-api"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	l     *log.Logger
	store product.Store
}

func NewProductHandler(store product.Store, l *log.Logger) *ProductHandler {
	return &ProductHandler{
		store: store,
		l:     l,
	}
}

func (ph *ProductHandler) Routes(r chi.Router) {
	r.Get("/", ph.GetAll)
	r.Post("/", ph.Post)

	r.Get("/{id}", ph.Get)
	r.Put("/{id}", ph.Put)
	r.Delete("/{id}", ph.Delete)
}

func (ph *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp := &response{}
	pl, err := ph.store.FindAll(context.TODO())

	produceError(
		w,
		500,
		err,
		"there is a problem retrieving records. Please try again later",
		"unable to retrieve products from store",
	)

	resp.Status = 200
	resp.Products = pl
	resp.sendJSON(w)
}

func (ph *ProductHandler) Post(w http.ResponseWriter, r *http.Request) {
	resp := &response{}
	p, err := parseBody(r)
	produceError(
		w,
		400,
		err,
		"invalid data, please refer to documentation for correct usage",
		"error parsing data",
	)

	p, err = ph.store.Create(context.TODO(), p)

	produceError(
		w,
		500,
		err,
		"error creating product, please try again later",
		"error creating product",
	)

	resp.Status = 201
	resp.Product = p
	resp.sendJSON(w)
}

func (ph *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	resp := &response{}
	id, err := parseId(r)

	produceError(
		w,
		400,
		err,
		"invalid id, please refer to documentation for correct usage",
		"received invalid id",
	)

	p, err := ph.store.FindById(context.TODO(), id)

	produceError(
		w,
		400,
		err,
		"something went wrong",
		"error retrieving record",
	)

	if p == nil {
		resp.Status = 404
		resp.Error = errors.New("Product with id " + strconv.Itoa(id) + "not found")
		resp.sendJSON(w)
		return
	}

	resp.Status = 200
	resp.Product = p
	resp.sendJSON(w)
}

func (ph *ProductHandler) Put(w http.ResponseWriter, r *http.Request) {
	resp := &response{}
	id, err := parseId(r)
	produceError(
		w,
		400,
		err,
		"invalid id",
		"invalid id",
	)
	p, err := parseBody(r)
	produceError(
		w,
		400,
		err,
		"invalid data",
		"invalid data",
	)

	p, err = ph.store.Update(context.TODO(), id, p)

	produceError(
		w,
		500,
		err,
		"error updating value",
		"error updating value",
	)

	if p == nil {
		resp.Status = 404
		resp.Error = errors.New("Product with id " + strconv.Itoa(id) + "not found")
		resp.sendJSON(w)
		return
	}

	resp.Status = 200
	resp.Product = p
	resp.sendJSON(w)
}

func (ph *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resp := &response{}
	id, err := parseId(r)
	produceError(
		w,
		400,
		err,
		"invalid id",
		"invalid id",
	)

	p, err := ph.store.Delete(context.TODO(), id)

	produceError(
		w,
		500,
		err,
		"error deleting value",
		"error deleting value",
	)

	if p == nil {
		resp.Status = 404
		resp.Error = errors.New("Product with id " + strconv.Itoa(id) + "not found")
		resp.sendJSON(w)
		return
	}

	resp.Status = 200
	resp.Product = p
	resp.sendJSON(w)
}

func parseBody(r *http.Request) (*product.Product, error) {
	p := product.New()
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		return nil, fmt.Errorf("unable to decode request body: %w", err)
	}

	return p, nil
}

func parseId(r *http.Request) (int, error) {
	idRaw := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		return -1, fmt.Errorf("id is not an integer: %w", err)
	}

	return id, nil
}

func produceError(w http.ResponseWriter, status int, err error, messageE string, messageI string) {
	if err != nil {
		resp := &response{}
		resp.Status = 400
		resp.Error = errors.New("something went wrong")
		resp.sendJSON(w)
		log.Panicln(fmt.Errorf("error retrieving record: %w", err))
	}
}
