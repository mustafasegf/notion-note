package api

import (
	"net/http"

	"github.com/mustafasegf/notion-note/controller"
	"github.com/mustafasegf/notion-note/repo"
	"github.com/mustafasegf/notion-note/service"
)

func (s *Server) SetupRouter() {
	lineRepo := repo.NewLineRepo(s.Db)
	lineService := service.NewLineService(s.Line, lineRepo)
	lineController := controller.NewLineController(s.Line, lineService)
	http.HandleFunc("/callback", lineController.LineCallback)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("yes"))
	})
}
