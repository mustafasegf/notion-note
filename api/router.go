package api

import (
	"net/http"

	"github.com/mustafasegf/notion-note/notes"
)

func (s *Server) SetupRouter() {
	lineRepo := notes.NewRepo(s.Db)
	lineService := notes.NewService(s.Line, lineRepo)
	lineController := notes.NewController(s.Line, lineService)

	http.HandleFunc("/callback", lineController.LineCallback)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("health check"))
	})
}
