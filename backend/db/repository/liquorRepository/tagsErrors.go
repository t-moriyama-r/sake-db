package liquorRepository

import (
	"backend/middlewares/customError"
	"backend/middlewares/customError/errorMsg"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

const (
	GetTags           = "REPO-LIQUOR-TAG-001-GetTags"
	GetTagsDecode     = "REPO-LIQUOR-TAG-002-GetTagsDecode"
	PostTag           = "REPO-LIQUOR-TAG-003-PostTag"
	DeleteTag         = "REPO-LIQUOR-TAG-004-DeleteTag"
	ZeroDelete        = "REPO-LIQUOR-TAG-005-ZeroDelete"
	SearchByTag       = "REPO-LIQUOR-TAG-006-SearchByTag"
	SearchByTagDecode = "REPO-LIQUOR-TAG-007-SearchByTagDecode"
)

func errGetTags(err error, liquorId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetTags,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.ErrorLevel,
		Input:      liquorId,
	})
}

func errGetTagsDecode(err error, liquorId primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    GetTagsDecode,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      liquorId,
	})
}

func errPostTag(err error, m *TagModel) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    PostTag,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      m,
	})
}

func errDeleteTag(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    DeleteTag,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      id,
	})
}
func errZeroDelete(err error, id primitive.ObjectID) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    ZeroDelete,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.InfoLevel,
		Input:      id,
	})
}

func errSearchByTag(err error, tag string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    SearchByTag,
		UserMsg:    errorMsg.DATA,
		Level:      logrus.ErrorLevel,
		Input:      tag,
	})
}

func errSearchByTagDecode(err error, tag string) *customError.Error {
	return customError.NewError(err, customError.Params{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    SearchByTagDecode,
		UserMsg:    errorMsg.SERVER,
		Level:      logrus.ErrorLevel,
		Input:      tag,
	})
}
