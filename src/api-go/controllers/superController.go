package controllers

import (
	"encoding/json"
	"net/http"
	"api-go/types"
	"fmt"
)

type SuperHeroCtrl interface {
	AddSuper(w http.ResponseWriter, r *http.Request)
	RemoveSuper(w http.ResponseWriter, r *http.Request)
	ListSuper(w http.ResponseWriter, r *http.Request)
	FindSuperById(w http.ResponseWriter, r *http.Request)
	FindSuperByName(w http.ResponseWriter, r *http.Request)
}

type superHeroCtrlImpl struct {
	srv types.SuperHeroService
}

func NewSuperHeroCtrl(srv types.SuperHeroService) SuperHeroCtrl {
	return &superHeroCtrlImpl{srv}
}

func (ctrl *superHeroCtrlImpl) AddSuper(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	enc := json.NewEncoder(w)
	fmt.Print(params)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not process request"})
	}

	err = ctrl.srv.AddSuper(params["super-name"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not save super"})
	}

	w.WriteHeader(200)
}

func (ctrl *superHeroCtrlImpl) RemoveSuper(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	enc := json.NewEncoder(w)
	fmt.Print(params)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not process request"})
	}

	err = ctrl.srv.RemoveSuper(params["super-name"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not erase super"})
	}

	w.WriteHeader(200)
}

func (ctrl *superHeroCtrlImpl) ListSuper(w http.ResponseWriter, r *http.Request)  {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	enc := json.NewEncoder(w)
	fmt.Print(params)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not find any super"})
	}

	var retorno []types.Get

	retorno = ctrl.srv.ListSuper(params["type"])

	if retorno == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not deliver response"})
	}
	json.NewEncoder(w).Encode(retorno)
}

func (ctrl *superHeroCtrlImpl) FindSuperByName(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	enc := json.NewEncoder(w)
	fmt.Print(params)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not find any super"})
	}

	var retorno []types.Get

	retorno = ctrl.srv.FindSuperByName(params["super-name"])
	if retorno == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not deliver response"})
	}
	json.NewEncoder(w).Encode(retorno)

}

func (ctrl *superHeroCtrlImpl) FindSuperById(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	err := json.NewDecoder(r.Body).Decode(&params)
	enc := json.NewEncoder(w)
	fmt.Print(params)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		enc.Encode(&map[string]string{"message": "Could not find any super"})
	}
	var retorno []types.Get

	retorno = ctrl.srv.FindSuperById(params["id"])

	if retorno == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		enc.Encode(&map[string]string{"message": "Could not deliver response"})
	}
	json.NewEncoder(w).Encode(retorno)

}
