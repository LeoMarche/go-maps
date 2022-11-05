package gomapsapi

import (
	"fmt"
	"net/http"
)

func (ws *WorkingSet) SolvePath(w http.ResponseWriter, r *http.Request) {

	// Check that the request is a POST
	if r.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Try to parse the form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		http.Error(w, "Please provide requested infos in form", http.StatusBadRequest)
		return
	}

	// TODO: check that form contains 'fromX', 'fromY' and 'toX' and 'toY' keys, compute the shortest path and return it

}
