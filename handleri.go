package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) DodavanjeKonfiga(w http.ResponseWriter, r *http.Request) {
	var config Config
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	s.data[id] = []*Config{&config}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) DodavanjeGrupe(w http.ResponseWriter, r *http.Request) {
	var config Config
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	s.data[id] = []*Config{&config}

	w.WriteHeader(http.StatusCreated)
}

func (s *Service) GetKonfiga(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	configs, ok := s.data[id]
	if !ok {
		http.Error(w, "config not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(configs)
}


func (s *Service) GetGrupe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	configs, ok := s.data[id]
	if !ok {
		http.Error(w, "config group not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(configs)
}



func (s *Service) BrisanjeKonfiga(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if _, ok := s.data[id]; !ok {
		http.Error(w, "config not found", http.StatusNotFound)
		return
	}

	delete(s.data, id)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) BrisanjeGrupe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if _, ok := s.data[id]; !ok {
		http.Error(w, "config group not found", http.StatusNotFound)
		return
	}

	delete(s.data, id)
	w.WriteHeader(http.StatusNoContent)
}


func (s *Service) DodavanjeKonfigaUGrupu(w http.ResponseWriter, r *http.Request) {
	var config Config
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	configs, ok := s.data[id]
	if !ok {
		http.Error(w, "config group not found", http.StatusNotFound)
		return
	}

	configs = append(configs, &config)
	s.data[id] = configs

	w.WriteHeader(http.StatusCreated)
}