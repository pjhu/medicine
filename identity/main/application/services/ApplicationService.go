package identityapplicationservice

import (
	"strings"
	cache "medicine/common/main/cache"
	cqrs "medicine/common/main/datasource"
	IdGenerator "medicine/common/main/idgenerator"
	identitycommand "medicine/identity/main/application/command"
	identityresponse "medicine/identity/main/application/response"
	identitymodel "medicine/identity/main/domain/models"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

// Signin for user register
func Signin(signinCommand identitycommand.SignoutCommand)  (response identityresponse.SigninResponse, e error) {
	has, err := cqrs.Engine.Exist(&identitymodel.Member{
		Phone: signinCommand.Phone,
	})
	if err != nil {
		log.Error(err)
		return response, err
	}
	var userId int64
	if !has {	
		session := cqrs.Engine.NewSession()
		defer session.Close()

		// add Begin() before any action
		err = session.Begin()
		if err != nil {
			log.Error(err)
			return response, err
		}
		userId = IdGenerator.NewId()
		newMember := identitymodel.NewMember(userId, signinCommand.Phone)
		_, err = cqrs.Engine.InsertOne(newMember)
		if err != nil {
			log.Error(err)
			session.Rollback()
			return response, err
		}
	} 
	
	token := strings.Replace(uuid.New().String(), "-", "", -1)
	userMeta := cache.UserMeta {Id: userId, Phone : signinCommand.Phone}
	cache.Set(cache.UserAuthNameSpace, token, userMeta)
	if err != nil {
		log.Error(err)
		return response, err
	}
	
	response = identityresponse.SigninResponse{Token: token}
	return response, err
}

// Signout for user logout
func Signout(token string)  {
	if (strings.HasPrefix(token, "MP")) {
		cache.Delete(cache.UserAuthNameSpace, token[3:])
	}
}