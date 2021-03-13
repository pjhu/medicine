package identityapplicationservice

import (
	log "github.com/sirupsen/logrus"
	
	"usercenter/common/main/errors"
	cache "usercenter/common/main/cache"
	cqrs "usercenter/common/main/datasource"
	IdGenerator "usercenter/common/main/idgenerator"
	identitycommand "usercenter/identity/main/application/command"
	identityresponse "usercenter/identity/main/application/response"
	identitymodel "usercenter/identity/main/domain/models"
)

// Signin for user register
func Signin(signinCommand identitycommand.SignoutCommand) (response identityresponse.SigninResponse, e *errors.ErrorWithCode) {
	has, err := cqrs.Engine.Exist(&identitymodel.Member{Phone: signinCommand.Phone,})
	if err != nil {
		log.Error(err)
		return response, errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
	}

	var userId int64
	if !has {	
		session := cqrs.Engine.NewSession()
		defer session.Close()
		// add Begin() before any action
		err = session.Begin()
		if err != nil {
			log.Error(err)
			return response, errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
		}
		userId = IdGenerator.NewId()
		newMember := identitymodel.NewMember(userId, signinCommand.Phone)
		_, err = cqrs.Engine.InsertOne(newMember)
		if err != nil {
			log.Error(err)
			session.Rollback()
			return response, errors.NewErrorWithCode(errors.SystemInternalError, "insert user error")
		}
	} 
	
	token := cache.CreateTokenKey()
	userMeta := cache.UserMeta {Id: userId, Phone : signinCommand.Phone}
	cache.StoreBy(cache.UserAuthNameSpace, token, userMeta)
	if err != nil {
		log.Error(err)
		return response, errors.NewErrorWithCode(errors.SystemInternalError, "cache user token error")
	}

	response = identityresponse.SigninResponse{Token: token}
	return response, nil
}

// Signout for user logout
func Signout(fullTokenString string) (e *errors.ErrorWithCode) {
	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		log.Error(err)
		return errors.NewErrorWithCode(errors.SystemInternalError, "error token")
	}
	err = cache.DeleteBy(cache.UserAuthNameSpace, tokenString)
	if err != nil {
		log.Error(err)
		return errors.NewErrorWithCode(errors.SystemInternalError, "delete token error")
	}
	return nil
}

// ValidateToken for validate token, refresh token, return user meta
func ValidateToken(fullTokenString string) (userMeta *cache.UserMeta, e *errors.ErrorWithCode) {
	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		log.Error(err)
		return userMeta, errors.NewErrorWithCode(errors.SystemInternalError, "error token format")
	}
	
	err = cache.GetBy(cache.UserAuthNameSpace, tokenString, &userMeta)
	if err != nil {
		log.Error(err)
		return userMeta, errors.NewErrorWithCode(errors.SystemInternalError, "token invalid")
	}
	return userMeta, nil
}

func refreshToken(token string) {
	log.Info("refresh token")
}