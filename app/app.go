package app

import (
	"expenditure-manager/src/repo"
	"fmt"
	"time"

	"github.com/sarulabs/di"
	mgo "gopkg.in/mgo.v2"
)

// Init inits app
func Init() di.Context {
	builder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}
	builder.AddDefinition(di.Definition{
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctx di.Context) (interface{}, error) {
			return mgo.DialWithInfo(&mgo.DialInfo{
				Addrs:   []string{"localhost:27017"},
				Timeout: time.Second * 30,
			})
		},
	})
	builder.AddDefinition(di.Definition{
		Name:  "expenditure-repo-factory",
		Scope: di.Request,
		Build: func(ctx di.Context) (interface{}, error) {
			fmt.Println("get rpeo dac")
			return &repo.ExpenditureRepoFactory{
				Session:    ctx.Get("mongo").(*mgo.Session),
				DataBase:   "moneymind",
				Collection: "expenditures",
			}, nil
		},
	})

	return builder.Build()
}
