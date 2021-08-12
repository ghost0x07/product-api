package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"product-api"
)

type response struct {
	Status   int                `json:"-"`
	Headers  map[string]string  `json:"-"`
	Product  *product.Product   `json:"product,omitempty"`
	Products []*product.Product `json:"products,omitempty"`
	Error    error              `json:"error,omitempty"`
}

func (resp *response) sendJSON(w http.ResponseWriter) {
	w.WriteHeader(resp.Status)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Panic(fmt.Errorf("error encoding %#v to JSON: %w", resp, err))
	}

}
