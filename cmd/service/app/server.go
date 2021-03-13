package app

import (
	"encoding/json"
	"fmt"
	"github.com/ArtemBond13/ago2_rest/pkg/offers"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Server struct {
	offersSvc *offers.Service
	router    chi.Router
}

func NewServer(offers *offers.Service, router chi.Router) {
	return &Server{offersSvc: offersSvc, router: router}
}

// настройка роутинга
func (s *Server) Init() error {
	s.router.Get("/offers", s.handleGetOffers)
	s.router.Get("/offers/{id}", s.handleGetOffersByID)
	s.router.Post("/offers", s.handleSaveOffer)
	s.router.Delete("/offers/{id}", s.handleRemoveOfferByID)

	return nil
}

func (s *Server) ServerHTTP(writer http.ResponseWriter, reader *http.Request) {
	s.router.ServeHTTP(writer, reader)
}

// у Get нет тела (только path, query и заголовки)
func (s *Server) handleGetOffers(writer http.ResponseWriter, request *http.Request) {
	items, err := s.offersSvc.All(request.Context())
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// серилизуем
	data, err := json.Marshal(items)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) handleGetOffersByID(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("not implemented")
}

func (s *Server) handleSaveOffer(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("not implemented")
}

func (s *Server) handleRemoveOfferByID(writer http.ResponseWriter, reader *http.Request) {
	fmt.Println("not implemented")
}
