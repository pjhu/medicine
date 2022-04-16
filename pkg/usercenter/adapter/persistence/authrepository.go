package persistence

import (
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"pjhu/medicine/pkg/usercenter/domain"
)

type Repo struct {
	DB *xorm.EngineGroup
}

func BuildMysqlRepo(db *xorm.EngineGroup) domain.IRepository {
	return &Repo{
		DB: db,
	}
}

func (r *Repo) InsertOne(member *domain.Member) (int64, error) {

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

	_, err = r.DB.InsertOne(member)
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

func (r *Repo) Get(member *domain.Member) (*domain.Member, error) {
	_, err := r.DB.Get(member)
	if err != nil {
		return nil, err
	}
	return member, err
}

func (r *Repo) Exist(member *domain.Member) (bool, error) {
	exist, err := r.DB.Exist(member)
	if err != nil {
		return false, err
	}
	return exist, err
}
