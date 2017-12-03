package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/n3wscott/hue-lights/pkg/apis/huebroker/v1"
	"github.com/n3wscott/hue-lights/pkg/controller"
	h1 "github.com/n3wscott/hue-lights/pkg/hue"
)

func (s *server) GetPeople(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	people, _ := s.Controller.GetPeople()
	json.NewEncoder(w).Encode(people)
}

func (s *server) GetPerson(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	if p, err := s.Controller.GetPerson(params["id"]); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

func (s *server) CreatePerson(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	var personRequest v1.CreatePersonRequest
	_ = json.NewDecoder(req.Body).Decode(&personRequest)

	if p, err := s.Controller.CreatePerson(params["id"], &personRequest); err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

func (s *server) DeletePerson(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	if err := s.Controller.DeletePerson(params["id"]); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *server) GetBridge(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	people, _ := s.Controller.GetBridge()
	json.NewEncoder(w).Encode(people)
}

type server struct {
	Router     *mux.Router
	Controller v1.Controller
}

func CreateServer() *server {

	people := []v1.Person{}
	people = append(people, v1.Person{ID: "1", FirstName: "Scott", LastName: "Nichols", Address: &v1.Address{City: "Seattle", State: "WA"}})
	people = append(people, v1.Person{ID: "2", FirstName: "Costin", LastName: "Nichols"})

	hue := h1.Hue{Bridge: &h1.Bridge{}}

	s := server{
		Router: mux.NewRouter(),
		Controller: &controller.PeopleController{
			People: people,
			Hue:    &hue,
		},
	}

	s.Router.HandleFunc("/people", s.GetPeople).Methods("GET")
	s.Router.HandleFunc("/people/{id}", s.GetPerson).Methods("GET")
	s.Router.HandleFunc("/people/{id}", s.CreatePerson).Methods("POST")
	s.Router.HandleFunc("/people/{id}", s.DeletePerson).Methods("DELETE")

	s.Router.HandleFunc("/bridge", s.GetBridge).Methods("GET")

	return &s
}
