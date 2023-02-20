package server

import (
	"net/http"
	"netforemost/utils/responses"
)
func  RoutesHandlers(mux *http.ServeMux){
	mux.HandleFunc("/notes", HandleNotes)
}

func HandleNotes(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			message := struct {
				Text string `json:"message"`
			}{
				Text: "das",
			}
			responses.JSON(w, http.StatusCreated, message)
		case http.MethodPost:
			// Post notes
		case http.MethodPut:
			// Put notes
		case http.MethodDelete:
			// Delete notes
		default:
			// Unknown method
	}
}