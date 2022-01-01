package api

import (
	"net/http"

	"github.com/mustafasegf/notion-note/notes"
)

func (s *Server) SetupRouter() {
	lineRepo := notes.NewRepo(s.Db)
	lineService := notes.NewService(s.Line, lineRepo)
	lineController := notes.NewController(s.Line, lineService)

	http.HandleFunc("/callback/line", lineController.LineCallback)
	http.Handle("/", http.FileServer(http.Dir("./templates")))
}
