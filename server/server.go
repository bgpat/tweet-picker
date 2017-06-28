package server

import (
	"os"

	"github.com/bgpat/tweet-picker/server/router"
	"gopkg.in/gin-gonic/gin.v1"
)

type Server struct {
	*gin.Engine
}

func New() *Server {
	r := gin.Default()
	router.Initialize(r)
	return &Server{
		Engine: r,
	}
}

func (s *Server) Run() error {
	socket := os.Getenv("UNIX_SOCKET")
	if socket != "" {
		return s.Engine.RunUnix(socket)
	}
	addr := os.Getenv("ADDR")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return s.Engine.Run(addr + ":" + port)
}
