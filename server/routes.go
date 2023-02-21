package server

import (
	"net/http"
	"netforemost/controllers"
)
func  RoutesHandlers(mux *http.ServeMux){
	mux.HandleFunc("/notes", HandleNotes)
}

func HandleNotes(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			controllers.GetNotes(Server.DB, w, r)
		case http.MethodPost:
			controllers.MyServer.CreateNote(Server.DB, w, r)
		case http.MethodPut:
			controllers.MyServer.UpdateNote(Server.DB, w, r)
		case http.MethodDelete:
			controllers.MyServer.DeleteNote(Server.DB, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}