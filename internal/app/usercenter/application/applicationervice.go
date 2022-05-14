package application

import (
	Error "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/pjhu/medicine/internal/app/usercenter/domain"
	"github.com/pjhu/medicine/internal/pkg/cache"
	"github.com/pjhu/medicine/internal/pkg/datasource"
	"github.com/pjhu/medicine/pkg/errors"
	"github.com/pjhu/medicine/pkg/idgenerator"
)

type IAuthApplicationService interface {
	Signin(signinCommand SigninCommand) (response SigninResponse, e *errors.ErrorWithCode)
	Signout(fullTokenString string) (e *errors.ErrorWithCode)
	ValidateToken(fullTokenString string) (userMeta *cache.UserMeta, e *errors.ErrorWithCode)
}

type AuthApplicationService struct {
	db    *gorm.DB
	repo  domain.IAuthRepository
	cache cache.ICacheRepository
}

// 变量断言: 如果没有实现所以方法，编译报missing method
var _ IAuthApplicationService = &AuthApplicationService{}

func NewAuthApplicationService(db *gorm.DB, repo domain.IAuthRepository,
	cache cache.ICacheRepository) *AuthApplicationService {
	return &AuthApplicationService{
		db:    db,
		repo:  repo,
		cache: cache,
	}
}

// Signin for user register
func (appSvc *AuthApplicationService) Signin(signinCommand SigninCommand) (result SigninResponse, e *errors.ErrorWithCode) {

	db := datasource.NewDBSession()
	exist := false
	userId := idgenerator.NewId()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := appSvc.repo.FindBy(&domain.Member{Phone: signinCommand.Phone})
		if err == nil {
			exist = true
		}
		if err != nil && Error.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error(err)
			newMember := domain.NewMember(userId, signinCommand.Phone)
			err = appSvc.repo.InsertOne(newMember)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	})

	if exist {
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "user phone already signin")
	}
	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "Signin failed")
	}

	token := cache.CreateTokenKey()
	userMeta := cache.UserMeta{Id: userId, Phone: signinCommand.Phone}
	err = appSvc.cache.StoreBy(cache.UserAuthNameSpace, token, userMeta)
	if err != nil {
		logrus.Error(err)
		return result, errors.NewErrorWithCode(errors.SystemInternalError, "cache user token error")
	}
	result = SigninResponse{Token: token}
	return result, nil
}

// Signout for user logout
func (appSvc *AuthApplicationService) Signout(fullTokenString string) (e *errors.ErrorWithCode) {
	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		logrus.Error(err)
		return errors.NewErrorWithCode(errors.SystemInternalError, "error token")
	}
	err = appSvc.cache.DeleteBy(cache.UserAuthNameSpace, tokenString)
	if err != nil {
		logrus.Error(err)
		return errors.NewErrorWithCode(errors.SystemInternalError, "delete token error")
	}
	return nil
}

// ValidateToken for validate token, refresh token, return user meta
func (appSvc *AuthApplicationService) ValidateToken(fullTokenString string) (userMeta *cache.UserMeta, e *errors.ErrorWithCode) {
	tokenString, err := cache.ExtractTokenKey(fullTokenString)
	if err != nil {
		logrus.Error(err)
		return userMeta, errors.NewErrorWithCode(errors.SystemInternalError, "error token format")
	}

	err = appSvc.cache.GetBy(cache.UserAuthNameSpace, tokenString, &userMeta)
	if err != nil {
		logrus.Error(err)
		return userMeta, errors.NewErrorWithCode(errors.SystemInternalError, "token invalid")
	}
	return userMeta, nil
}

func refreshToken(token string) {
	logrus.Info("refresh token")
}
