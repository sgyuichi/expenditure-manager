package repo

import (
	entity "expenditure-manager/src/entity"

	mgo "gopkg.in/mgo.v2"
)

// ExpenditureRepoFactory -
type ExpenditureRepoFactory struct {
	Session    *mgo.Session
	DataBase   string
	Collection string
}

// Build builds an ExpenditureRepo
func (fac *ExpenditureRepoFactory) Build() *ExpenditureRepo {
	return &ExpenditureRepo{
		Collection: fac.Session.DB(fac.DataBase).C(fac.Collection),
	}
}

// ExpenditureRepo -
type ExpenditureRepo struct {
	Collection *mgo.Collection
}

// FindByID returns an Expenditure given its id
func (r *ExpenditureRepo) FindByID(id string) (*entity.Expenditure, error) {
	var res *entity.Expenditure
	err := r.Collection.FindId(id).One(&res)
	return res, err
}

// Insert inserts an Expenditure
func (r *ExpenditureRepo) Insert(e *entity.Expenditure) error {
	return r.Collection.Insert(e)
}
