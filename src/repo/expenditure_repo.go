package repo

import (
	entity "expenditure-manager/src/entity"

	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// FindBetween returns an Expenditure given its id
func (r *ExpenditureRepo) FindBetween(from, to int64) ([]*entity.Expenditure, error) {
	res := []*entity.Expenditure{}
	err := r.Collection.Find(bson.M{"date": bson.M{"$gte": from, "$lte": to}}).All(&res)
	return res, err
}

// Insert inserts an Expenditure
func (r *ExpenditureRepo) Insert(e *entity.Expenditure) error {
	if e.ID == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}
		e.ID = id.String()
	}
	return r.Collection.Insert(e)
}
