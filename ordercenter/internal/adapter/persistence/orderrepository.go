package persistence

import (
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"ordercenter/internal/domain"
)

type Repo struct {
	DB *xorm.EngineGroup
}

func BuildMysqlRepo(db *xorm.EngineGroup) domain.IRepository {
	return &Repo {
		DB: db,
	}
}

func (r *Repo) InsertOne(userOrder *domain.UserOrder) (int64, error) {

	session := r.DB.NewSession()
	err := session.Close()
	if err != nil {
		return 0, err
	}

	// add Begin() before any action
	err = session.Begin()
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	_, err = r.DB.InsertOne(userOrder)
	if err != nil {
		logrus.Error(err)
		err := session.Rollback()
		if err != nil {
			logrus.Error(err)
		}
		return 0, err
	}
	return 1, err
}

func (r *Repo) Get(userOrder *domain.UserOrder) (*domain.UserOrder, error) {
	_, err := r.DB.Get(userOrder)
	if err != nil {
		return nil, err
	}
	return userOrder, err
}