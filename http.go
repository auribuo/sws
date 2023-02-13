package main

import (
	"fmt"
	"net/http"
)

type ListenAndServer interface {
	ListenAndServe() error
}

type HttpServer struct {
	Dir  string
	Port int
}

func (s *HttpServer) ListenAndServe() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: http.FileServer(http.Dir(s.Dir)),
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
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: http.FileServer(http.Dir(s.Dir)),
	}
	return server.ListenAndServeTLS(s.Cert, s.Key)
}

func NewServer(dir string, port int, https bool, cert string, key string) ListenAndServer {
	if https {
		return &HttpsServer{Dir: dir, Port: port, Cert: cert, Key: key}
	}
	return &HttpServer{Dir: dir, Port: port}
}
