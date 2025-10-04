package main

import (
	// "encoding/json"
	"log"
	"net/http"
)

type BrokerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	app := Config{}

	http.HandleFunc("/broker", app.Broker)

	log.Println("Broker service listening on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// func brokerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	resp := BrokerResponse{
// 		Status:  "success",
// 		Message: "Broker service is alive!",
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp)
// }
