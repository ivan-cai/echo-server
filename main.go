package main

import (
	"encoding/json"
        "fmt"
	logs "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
        fmt.Prinln("add log for test")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs.Infof("Hello Echo Server")
		json.NewEncoder(w).Encode("Hello Echo Server")
	})

	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		logs.Infof("Hello Echo Server v2.0 @oversea")
		json.NewEncoder(w).Encode("Hello Echo Server v2.0 @oversea")
	})

	http.ListenAndServe(":8080", router)
}
