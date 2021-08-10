package api

import (
	"net/http"

	"github.com/mustafasegf/notion-note/controller"
	"github.com/mustafasegf/notion-note/service"
)

func (s *Server) SetupRouter() {

	lineService := service.NewLinkService(s.Line, s.Notion)
	lineController := controller.NewLinkController(s.Line, lineService)
	http.HandleFunc("/callback", lineController.LineCallback)
}
