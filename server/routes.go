package server

import (
	"net/http"
	"netforemost/controllers"
)
func (s *Server)  RoutesHandlers(mux *http.ServeMux){
	mux.HandleFunc("/notes", s.HandleNotes)
}

func (s *Server) HandleNotes(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			controllers.GetNotes(s.DB, w, r)
		case http.MethodPost:
			controllers.CreateNote(s.DB, w, r)
		case http.MethodPut:
			controllers.UpdateNote(s.DB, w, r)
		case http.MethodDelete:
			controllers.DeleteNote(s.DB, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}