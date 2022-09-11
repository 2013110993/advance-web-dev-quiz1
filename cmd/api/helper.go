// Filename: cmd/api/helpers.go

package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Function for reading ID
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the param
	parameter := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(parameter.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid ID ")
	}

	return id, nil

}
