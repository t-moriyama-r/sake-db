package auth

import (
	"backend/middlewares/customError"
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	InvalidToken   = "AUTH-001-InvalidToken"
	ExpireToken    = "AUTH-002-ExpireToken"
	BugToken       = "AUTH-003-BugToken"
	NotFoundToken  = "AUTH-004-NotFoundToken"
	NotFoundBearer = "AUTH-005-NotFoundBearer"
	NotFoundUser   = "AUTH-006-NotFoundUser"
	UnAuthorized   = "AUTH-007-UnAuthorized"
)

func errTokenInvalid(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    InvalidToken,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenExpired(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    ExpireToken,
		UserMsg:    "トークンが期限切れです。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenSomething(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    BugToken,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errMissHeader() *customError.Error {
	return customError.NewError(errors.New("authorizationヘッダーが見つかりません"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundToken,
		UserMsg:    "トークンが見つかりません",
		Level:      logrus.InfoLevel,
	})
}

func errMissBearer() *customError.Error {
	return customError.NewError(errors.New("authorizationトークンが見つかりません"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundBearer,
		UserMsg:    "トークンが見つかりません",
		Level:      logrus.InfoLevel,
	})
}

func errNotFoundUser() *customError.Error {
	return customError.NewError(errors.New("ユーザーIDがが見つかりません"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundUser,
		UserMsg:    "ユーザーIDがが見つかりません",
		Level:      logrus.InfoLevel,
	})
}
func errUnAuthorized(id primitive.ObjectID) *customError.Error {
	return customError.NewError(errors.New("ユーザーIDがが見つかりません"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    UnAuthorized,
		UserMsg:    "ユーザーIDがが見つかりません",
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}
