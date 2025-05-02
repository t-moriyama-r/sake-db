package authService

import (
	"backend/db/repository/userRepository"
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	TokenNotFound       = "TOKEN-001"
	TokenExpired        = "TOKEN-002"
	TokenInvalid        = "TOKEN-003"
	TokenInvalidClimes  = "TOKEN-004"
	RefreshTokenInvalid = "TOKEN-005"
)

const (
	NotFoundPassOrMail = "LOGIN-001"
)

func errTokenNotFound() *customError.Error {
	return customError.NewError(errors.New("refresh token not found"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenNotFound,
		UserMsg:    "トークンがありません。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenExpired() *customError.Error {
	return customError.NewError(errors.New("expired refresh token"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenExpired,
		UserMsg:    "トークンが期限切れです。",
		Level:      logrus.InfoLevel,
	})
}

func errTokenInvalid() *customError.Error {
	return customError.NewError(errors.New("invalid refresh token"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenInvalid,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errInvalidClimes() *customError.Error {
	return customError.NewError(errors.New("invalid claims"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    TokenInvalidClimes,
		UserMsg:    "トークンが不正です。",
		Level:      logrus.InfoLevel,
	})
}

func errRefreshToken(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    RefreshTokenInvalid,
		UserMsg:    "リフレッシュトークンが期限切れです。",
		Level:      logrus.InfoLevel,
	})
}

func errLogin() *customError.Error {
	return customError.NewError(errors.New("メールアドレスもしくはパスワードが間違っています。"), customError.Params{
		StatusCode: http.StatusUnauthorized,
		ErrCode:    NotFoundPassOrMail,
		UserMsg:    "メールアドレスもしくはパスワードが間違っています。",
		Level:      logrus.InfoLevel,
	})
}

const (
	GenerateFromPassword = "AUTH-PASSWORD-RESET-001-GenerateFromPassword"
	SendPasswordReset    = "AUTH-PASSWORD-RESET-002-SendPasswordReset"
	GenerateAccessToken  = "AUTH-PASSWORD-RESET-003-GenerateAccessToken"
	GenerateRefreshToken = "AUTH-PASSWORD-RESET-004-GenerateRefreshToken"
)

func errGenerateFromPassword(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GenerateFromPassword,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
	})
}

func errSendPasswordReset(err error, email string, token string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    SendPasswordReset,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      fmt.Sprintf("email: %s, token: %s", email, token),
	})
}

func errGenerateAccessToken(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GenerateAccessToken,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

func errGenerateRefreshToken(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GenerateRefreshToken,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}

const (
	NeedPassword  = "AUTH-Register-001-NeedPassword"
	FailHash      = "AUTH-Register-002-FailHash"
	DuplicateMail = "AUTH-Register-003-DuplicateMail"
	FailRegister  = "AUTH-Register-004-FailRegister"
)

func errNeedPassword() *customError.Error {
	return customError.NewError(errors.New("パスワードは必須です"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    NeedPassword,
		UserMsg:    "パスワードは必須です",
		Level:      logrus.InfoLevel,
	})
}
func errFailHash(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FailHash,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.FatalLevel,
	})
}

func errDuplicateMail(u userRepository.Model) *customError.Error {
	return customError.NewError(errors.New("このメールアドレスは既に登録されています。"), customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    DuplicateMail,
		UserMsg:    "このメールアドレスは既に登録されています。",
		Level:      logrus.InfoLevel,
		Input:      u,
	})
}

func errFailRegister(cErr *customError.Error, u userRepository.Model) *customError.Error {
	return customError.NewError(cErr, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    FailRegister,
		UserMsg:    "ユーザー登録に失敗しました。",
		Level:      logrus.ErrorLevel,
		Input:      u,
	})
}
