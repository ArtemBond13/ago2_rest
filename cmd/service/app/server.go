package server

import (
	"encoding/json"
	"github.com/ArtemBond13/ago2_rest/pkg/offers"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	offersSvc *offers.Service
	router    chi.Router
}

func NewServer(offersSvc *offers.Service, router chi.Router) *Server {
	return &Server{offersSvc: offersSvc, router: router}
}

// настройка роутинга
func (s *Server) Init() error {
	s.router.Get("/offers", s.handleGetOffer)
	s.router.Get("/offers/{id}", s.handleGetOffersByID)
	s.router.Post("/offers", s.handleSaveOffer)
	s.router.Delete("/offers/{id}", s.handleRemoveOfferByID)

	return nil
}

func (s *Server) ServerHTTP(writer http.ResponseWriter, reader *http.Request) {
	s.router.ServeHTTP(writer, reader)
}

// у Get нет тела (только path, query и заголовки)
func (s *Server) handleGetOffer(writer http.ResponseWriter, request *http.Request) {
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

func (s *Server) handleGetOffersByID(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.offersSvc.ByID(request.Context(), id)
	if err != nil {
		log.Println(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		log.Println(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) handleSaveOffer(writer http.ResponseWriter, request *http.Request) {
	itemToSave := &offers.Offer{}
	// когда данные приходят по сети эффктивнее использовать json.Decode, json.Unmarshal
	err := json.NewDecoder(request.Body).Decode(&itemToSave)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.offersSvc.Save(request.Context(), itemToSave)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
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

func (s *Server) handleRemoveOfferByID(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	item, err := s.offersSvc.Delete(request, id)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)
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
