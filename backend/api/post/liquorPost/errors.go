package liquorPost

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
)

const (
	ParseFailInput         = "LIQUOR-POST-001-ParseFailInput"
	ParseTempID            = "LIQUOR-POST-002-ParseTempID"
	DuplicateName          = "LIQUOR-POST-003-DuplicateName"
	ParseID                = "LIQUOR-POST-004-ParseID"
	InvalidInput           = "LIQUOR-POST-005-InvalidInput"
	InvalidVersion         = "LIQUOR-POST-006-InvalidVersion"
	InvalidFile            = "LIQUOR-POST-007-InvalidFile"
	FlavorMapCategoryMove  = "LIQUOR-POST-008-FlavorMapCategoryMove"
)

func errInvalidInput(c *gin.Context, err error) *customError.Error {
	raw, getRawErr := c.GetRawData()
	if getRawErr != nil {
		return customError.NewError(err, customError.Params{
			StatusCode: http.StatusBadRequest,
			ErrCode:    ParseFailInput,
			UserMsg:    "入力値が不正です",
			Level:      logrus.ErrorLevel,
		})
	}
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidInput,
		UserMsg:    "入力値が不正です",
		Level:      logrus.ErrorLevel,
		Input:      raw,
	})
}

func errParseTempID(err error) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ParseTempID,
		UserMsg:    "IDが不正です",
		Level:      logrus.InfoLevel,
	})
}
func errParseID(err error, id string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ParseID,
		UserMsg:    "IDが不正です",
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errDuplicateName() *customError.Error {
	return customError.NewError(errors.New("すでに存在するお酒です"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    DuplicateName,
		UserMsg:    "すでに存在するお酒です",
		Level:      logrus.InfoLevel,
	})
}
func errInvalidVersion() *customError.Error {
	return customError.NewError(errors.New(errorMsg.VERSION), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidVersion,
		UserMsg:    "データが更新されました。再度お試し下さい。",
		Level:      logrus.InfoLevel,
	})
}
func errInvalidFile(err error, img multipart.File) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidFile,
		UserMsg:    "画像の読み込みに失敗しました",
		Level:      logrus.ErrorLevel,
		Input:      img,
	})
}

func errFlavorMapCategoryMove(oldCategoryName string, newCategoryName string) *customError.Error {
	return customError.NewError(errors.New("異なるフレーバーマップカテゴリへの移動はできません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    FlavorMapCategoryMove,
		UserMsg:    "このお酒は「" + oldCategoryName + "」カテゴリに属しているため、「" + newCategoryName + "」カテゴリには移動できません。フレーバーマップの意味が異なるカテゴリへの移動は禁止されています。",
		Level:      logrus.InfoLevel,
	})
}
