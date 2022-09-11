// Filename: cmd/api/helpers.go

package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Function for reading ID or errors
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the param
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id ")
	}

	return id, nil

}
