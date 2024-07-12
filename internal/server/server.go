package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	conroller "github.com/snusEbjoer/cipherService/internal/controller"
	"github.com/snusEbjoer/cipherService/internal/redis"
)

type Server struct {
	router *http.ServeMux
	port   string
	host   string
	ctx    context.Context
}

func NewServer() *Server {
	s := &Server{
		router: http.NewServeMux(),
		port:   os.Getenv("SERVER_PORT"),
		host:   os.Getenv("SERVER_HOST"),
		ctx:    context.Background(),
	}
	s.initRouter()
	return s
}

func (s *Server) initRouter() {
	ctx := context.Background()
	e := conroller.NewController(redis.NewRedis(), ctx)

	s.router.HandleFunc("POST /sha256/{string}", e.EncodeSHA256())
	s.router.HandleFunc("POST /md5/{string}", e.EncodeMD5())
}

func (s *Server) Run() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.host, s.port), s.router))
}
