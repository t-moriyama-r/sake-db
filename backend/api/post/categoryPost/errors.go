package categoryPost

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	ParseFailInput       = "CATEGORY-POST-001-ParseFailInput"
	InvalidInput         = "CATEGORY-POST-002-InvalidInput"
	InvalidParent        = "CATEGORY-POST-003-InvalidParent"
	InvalidVersion       = "CATEGORY-POST-004-InvalidVersion"
	InvalidFile          = "CATEGORY-POST-005-InvalidFile"
	DuplicateName        = "CATEGORY-POST-006-DuplicateName"
	ParentCategoryMove   = "CATEGORY-POST-007-ParentCategoryMove"
	ReadonlyCategoryMove = "CATEGORY-POST-008-ReadonlyCategoryMove"
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

func errInvalidParent(input RequestData) *customError.Error {
	return customError.NewError(errors.New("自身または子カテゴリを親とすることはできません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidParent,
		UserMsg:    "自身または子カテゴリを親とすることはできません",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errInvalidVersion(input RequestData) *customError.Error {
	return customError.NewError(errors.New("自身または子カテゴリを親とすることはできません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidVersion,
		UserMsg:    errorMsg.VERSION,
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errInvalidFile(err error, input RequestData) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    InvalidFile,
		UserMsg:    "ファイルが不正です",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errDuplicateName(input RequestData) *customError.Error {
	return customError.NewError(errors.New("同じ親カテゴリ内に同名のカテゴリが既に存在します"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    DuplicateName,
		UserMsg:    "同じ親カテゴリ内に同名のカテゴリが既に存在します",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errParentCategoryMove(input RequestData) *customError.Error {
	return customError.NewError(errors.New("親カテゴリは移動できません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ParentCategoryMove,
		UserMsg:    "親カテゴリは移動できません",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}

func errReadonlyCategoryMove(input RequestData) *customError.Error {
	return customError.NewError(errors.New("このカテゴリは移動できません"), customError.Params{
		StatusCode: http.StatusBadRequest,
		ErrCode:    ReadonlyCategoryMove,
		UserMsg:    "このカテゴリは移動できません",
		Level:      logrus.InfoLevel,
		Input:      input,
	})
}
