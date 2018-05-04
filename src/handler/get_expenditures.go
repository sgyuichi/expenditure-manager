package handler

import (
	"encoding/json"
	"expenditure-manager/src/repo"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sarulabs/di"
)

// GetExpenditures gets expenditures
func GetExpenditures(w http.ResponseWriter, r *http.Request, ctx di.Context) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var request struct {
		From int64 `json:"from"`
		To   int64 `json:"to"`
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		panic(err)
	}
	expenditureRepoFactory := ctx.Get("expenditure-repo-factory").(*repo.ExpenditureRepoFactory)
	if expenditureRepoFactory == nil {
		panic(fmt.Errorf("repo factory is nil"))
	}
	expenditureRepo := *expenditureRepoFactory.Build()
	exps, err := expenditureRepo.FindBetween(request.From, request.To)
	if err != nil {
		panic(err)
	}
	resp, err := json.Marshal(map[string]interface{}{"expenditures": exps})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
