package server

import (
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	counter int
}

func (s *Server) HandleAddtoCount(w http.ResponseWriter, r *http.Request) {
	s.counter++
	log.Println("count:", s.counter)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) HandleGetCount(w http.ResponseWriter, r *http.Request) {
	log.Println("value of count:", s.counter)
	_, err := w.Write([]byte(strconv.Itoa(s.counter)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
