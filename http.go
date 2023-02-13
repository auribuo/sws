package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	}
}

type ListenAndServer interface {
	ListenAndServe() error
}

type HttpServer struct {
	Dir  string
	Port int
}

func (s *HttpServer) ListenAndServe() error {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(NoCache())
	router.Static("/", s.Dir)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: router,
	}
	return server.ListenAndServe()
}

type HttpsServer struct {
	Dir  string
	Port int
	Cert string
	Key  string
}

func (s *HttpsServer) ListenAndServe() error {
	router := gin.New()
	router.Static("/", s.Dir)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(NoCache())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: router,
	}
	return server.ListenAndServeTLS(s.Cert, s.Key)
}

func NewServer(dir string, port int, https bool, cert string, key string) ListenAndServer {
	if https {
		return &HttpsServer{Dir: dir, Port: port, Cert: cert, Key: key}
	}
	return &HttpServer{Dir: dir, Port: port}
}
