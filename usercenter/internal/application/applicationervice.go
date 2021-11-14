package service

import (
	"github.com/sirupsen/logrus"
	"usercenter/internal/application/command"
	"usercenter/internal/application/response"
	"usercenter/internal/domain"
	"usercenter/internal/pkg/cache"
	"usercenter/internal/pkg/errors"
	"usercenter/internal/pkg/idgenerator"
)

type IApplicationService interface {
	Signin(signinCommand command.SigninCommand) (response response.SigninResponse, e *errors.ErrorWithCode)
	Signout(fullTokenString string) (e *errors.ErrorWithCode)
	ValidateToken(fullTokenString string) (userMeta *cache.UserMeta, e *errors.ErrorWithCode)
}

type AuthApplicationService struct {
	repository domain.IRepository
	rdbRepository cache.ICacheRepository
}

func Build(repo domain.IRepository, rdbRepo cache.ICacheRepository) IApplicationService {
	return AuthApplicationService {
		repository: repo,
		rdbRepository: rdbRepo,
	}
}

// Signin for user register
func (appSvc AuthApplicationService) Signin(signinCommand command.SigninCommand) (result response.SigninResponse, e *errors.ErrorWithCode) {
	has, err := appSvc.repository.Exist(&domain.Member{Phone: signinCommand.Phone,})

	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, err.Error())
	}

	var userId int64
	if !has {
		userId = idgenerator.NewId()
		newMember := domain.NewMember(userId, signinCommand.Phone)
		_, err := appSvc.repository.InsertOne(newMember)
		if err != nil {
			return result, errors.NewErrorWithCode(errors.SystemInternalError, "insert user to db failure")
		}
	} 
	
	token := cache.CreateTokenKey()
	userMeta := cache.UserMeta {Id: userId, Phone : signinCommand.Phone}
	err = appSvc.rdbRepository.StoreBy(cache.UserAuthNameSpace, token, userMeta)
	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "cache user token error")
	}
	result = response.SigninResponse{Token: token}
	return result, nil
}

// Signout for user logout
func (appSvc AuthApplicationService) Signout(fullTokenString string) (e *errors.ErrorWithCode) {
	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		logrus.Error(err)
		return errors.NewErrorWithCode(errors.SystemInternalError, "error token")
	}
	err = appSvc.rdbRepository.DeleteBy(cache.UserAuthNameSpace, tokenString)
	if err != nil {
		logrus.Error(err)
		return errors.NewErrorWithCode(errors.SystemInternalError, "delete token error")
	}
	return nil
}

// ValidateToken for validate token, refresh token, return user meta
func (appSvc AuthApplicationService) ValidateToken(fullTokenString string) (userMeta *cache.UserMeta, e *errors.ErrorWithCode) {
	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		logrus.Error(err)
		return userMeta, errors.NewErrorWithCode(errors.SystemInternalError, "error token format")
	}
	
	err = appSvc.rdbRepository.GetBy(cache.UserAuthNameSpace, tokenString, &userMeta)
	if err != nil {
		logrus.Error(err)
		return userMeta, errors.NewErrorWithCode(errors.SystemInternalError, "token invalid")
	}
	return userMeta, nil
}

func refreshToken(token string) {
	logrus.Info("refresh token")
}