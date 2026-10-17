package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *HttpHandlers
}

func NewServer(h *HttpHandlers) *Server {
	return &Server{
		handlers: h,
	}
}

func (s *Server)LaunchServer() error {
	router := mux.NewRouter()

	router.Path("/task").Methods("POST").HandlerFunc(s.handlers.HandleCreateTask)
	router.Path("/task{name}").Methods("GET").HandlerFunc(s.handlers.HandleGetTask)
	router.Path("/task").Methods("GET").HandlerFunc(s.handlers.HandleGetAllTasks)
	router.Path("/task").Methods("GET").HandlerFunc(s.handlers.HandleGetAllTasks)
	router.Path("/task{name}").Methods("PATCH").Queries("completed", "true").HandlerFunc(s.handlers.HandleCompleteTask)
	router.Path("/task{name}").Methods("DELETE").HandlerFunc(s.handlers.HandleDeleteTask)
	router.Path("/task{name}").Methods("PATCH").HandlerFunc(s.handlers.HandleUpdateTask)

	return http.ListenAndServe(":9091", router)
}
