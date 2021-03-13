package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	offersSvc *offers.Service
	router chi.Router
}

func NewServer(offers *offers.Service, router chi.Router) {
	return &Server{offersSvc: offersSvc, router: router}
}

// настройка роутинга
func (s *Server) Init() error {
	s.router.Get("/offers", s.handleGetOffers)
	s.router.Get("/offers/{id}", s.handleGetOffersByID)
	s.router.Post("/offers", s.handleOffer)
	s.router.Delete("/offers/{id}", s.handleRemoveOfferByID)

	return nil
}

func (s *Server) ServerHTTP(writer http.ResponseWriter, reader *http.Request)  {
	s.router.ServeHTTP(writer, reader)
}

func (s *Server) handleGetOffers(writer http.ResponseWriter, reader *http.Request)  {
	fmt.Print("not implemented")
}

func (s *Server) handleGetOffersByID(writer http.ResponseWriter, reader *http.Request)  {
	fmt.Println("not implemented")
}

func (s *Server) handleSaveOffer(writer http.ResponseWriter, reader *http.Request)  {
	fmt.Println("not implemented")
}

func (s *Server) handleRemoveOfferByID(writer http.ResponseWriter, reader *http.Request)  {
	fmt.Println("not implemented")
}

