package gomapsapi

import (
	"encoding/json"
	"net/http"
)

func (ws *WorkingSet) Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
	rc := ReturnCode{Status: "OK"}
	js, err := json.Marshal(rc)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
