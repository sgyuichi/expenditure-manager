package main

import (
	"expenditure-manager/app"
	"expenditure-manager/src/handler"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	app := app.Init()
	defer app.Delete()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/expenditures", func(w http.ResponseWriter, r *http.Request) {
		request, err := app.SubContext()
		if err != nil {
			panic(err)
		}
		defer request.Delete()
		fmt.Println("post")
		handler.PostExpenditure(w, r, request)
	}).Methods("POST")
	log.Info("listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
