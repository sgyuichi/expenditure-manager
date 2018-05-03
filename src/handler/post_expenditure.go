package handler

import (
	"encoding/json"
	"expenditure-manager/src/entity"
	"expenditure-manager/src/repo"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sarulabs/di"
)

// PostExpenditure creates expenditure
func PostExpenditure(w http.ResponseWriter, r *http.Request, ctx di.Context) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var exp *entity.Expenditure
	err = json.Unmarshal(body, &exp)
	if err != nil {
		panic(err)
	}
	expenditureRepoFactory := ctx.Get("expenditure-repo-factory").(*repo.ExpenditureRepoFactory)
	if expenditureRepoFactory == nil {
		panic(fmt.Errorf("repo factory is nil"))
	}
	expenditureRepo := *expenditureRepoFactory.Build()
	err = expenditureRepo.Insert(exp)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(204)
	w.Header().Set("Content-Type", "application/json")
	w.Write(nil)
}
