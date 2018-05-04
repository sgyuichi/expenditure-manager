package main

import (
	"expenditure-manager/app"
	"expenditure-manager/src/handler"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	app := app.Init()
	defer app.Delete()

	request, err := app.SubContext()
	if err != nil {
		panic(err)
	}
	defer request.Delete()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/expenditures", func(w http.ResponseWriter, r *http.Request) {
		handler.PostExpenditure(w, r, request)
	}).Methods("POST")
	router.HandleFunc("/expenditures/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetExpenditure(w, r, request)
	}).Methods("GET")
	router.HandleFunc("/expenditures", func(w http.ResponseWriter, r *http.Request) {
		handler.GetExpenditures(w, r, request)
	}).Methods("GET")

	log.Info("listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
