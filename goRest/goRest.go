package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Router struct {
	id   string `json:"id`
	Name string `json:"name,omitempty`
	Vrf  *Vrf   `json:"vrf,omitempty"`
}
type Vrf struct {
	id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

var routers []Router

func GetRouterEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range routers {
		if item.id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Router{})
}
func GetRoutersEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(routers)
}
func CreateRouterEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var router Router
	_ = json.NewDecoder(r.Body).Decode(&router)
	router.id = params["id"]
	routers = append(routers, router)
	json.NewEncoder(w).Encode(routers)
}
func DeleteRouterEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range routers {
		if item.id == params["id"] {
			routers = append(routers[:index], routers[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(routers)
	}
}

func main() {
	httpRouter := mux.NewRouter()
	routers = append(routers, Router{id: "1", Name: "router-blue", Vrf: &Vrf{id: "1", Name: "Blue"}})
	routers = append(routers, Router{id: "2", Name: "router-red", Vrf: &Vrf{id: "2", Name: "Red"}})
	httpRouter.HandleFunc("/routers", GetRoutersEndpoint).Methods("GET")
	httpRouter.HandleFunc("/router/{id}", GetRouterEndpoint).Methods("GET")
	httpRouter.HandleFunc("/router/{id}", CreateRouterEndpoint).Methods("POST")
	httpRouter.HandleFunc("/router/{id}", DeleteRouterEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", httpRouter))
}
