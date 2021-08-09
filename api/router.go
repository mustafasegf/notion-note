package api

import (
	"net/http"

	"github.com/mustafasegf/notion-note/controller"
)

func (s *Server) SetupRouter() {

	lineController := controller.NewLinkController(s.Line, s.Notion)
	http.HandleFunc("/callback", lineController.LineCallback)
}
