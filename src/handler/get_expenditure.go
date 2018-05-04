package handler

import (
	"encoding/json"
	"expenditure-manager/src/repo"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

// GetExpenditure gets expenditure
func GetExpenditure(w http.ResponseWriter, r *http.Request, ctx di.Context) {
	vars := mux.Vars(r)
	id := vars["id"]
	expenditureRepoFactory := ctx.Get("expenditure-repo-factory").(*repo.ExpenditureRepoFactory)
	if expenditureRepoFactory == nil {
		panic(fmt.Errorf("repo factory is nil"))
	}
	expenditureRepo := *expenditureRepoFactory.Build()
	exp, err := expenditureRepo.FindByID(id)
	if err != nil {
		panic(err)
	}
	resp, err := json.Marshal(map[string]interface{}{"expenditure": exp})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
