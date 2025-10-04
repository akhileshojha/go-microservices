package main

import (
	"fmt"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.errorJSON(w, fmt.Errorf("Method Not Allowed"), http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Broker endpoint hit with method:", r.Method)
	payload := jsonResponse{
		Error:   false,
		Message: "Broker service is alive!",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}
