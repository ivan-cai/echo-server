package main

import (
	"encoding/json"
	logs "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//logs.Infof("Hello Echo Server for yunqi demo 1")
		echo()
		json.NewEncoder(w).Encode("Hello Echo Server for Demo 1")
	})

	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		logs.Infof("Hello Echo Server v2.0 @oversea")
		json.NewEncoder(w).Encode("Hello Echo Server v2.0 @oversea")
	})

	http.ListenAndServe(":8080", router)
}

func echo() {

}
