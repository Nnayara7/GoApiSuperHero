package routes

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	 "net/http"
	 "api-go/types"
)

func NewHeroRouter(ctrl types.SuperHeroCtrl) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/super", ctrl.AddSuper).Methods("POST")
	r.HandleFunc("/super", ctrl.ListSuper).Methods("GET")
	r.HandleFunc("/superName", ctrl.FindSuperByName).Methods("GET")
	r.HandleFunc("/superId", ctrl.FindSuperById).Methods("GET")
	r.HandleFunc("/super", ctrl.RemoveSuper).Methods("DELETE")
	return cors.Default().Handler(r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Welcome to Naiara API")
	w.Write(msg)
}


