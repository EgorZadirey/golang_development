package server

import (
	"encoding/json"
	. "finalWork/structures"
	"github.com/gorilla/mux"
	"net/http"
)

func handleConnection(w http.ResponseWriter, r *http.Request) {

	resultT := ResultT{}
	resultGet := GetResultData()
	resultT.Status = true
	resultT.Data = resultGet
	res1, _ := json.Marshal(resultT)
	_, err := w.Write(res1)
	if err != nil {
		return
	}
}

func RunApp() {
	server := http.Server{Addr: "localhost:8282"}
	router := mux.NewRouter()
	router.HandleFunc("/", handleConnection)
	http.Handle("/", router)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
