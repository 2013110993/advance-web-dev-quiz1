//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

//Entry Handler - Post
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create New Entry")
}

//Entry Handler - Get
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Show ENtry ID
	fmt.Fprintf(w, "Show Entry: %d\n", id)
}
