package auth

import "backend/middlewares/customError"

// IsErrWithoutAuth ユーザー認証エラーを除外する
func IsErrWithoutAuth(err *customError.Error) bool {
	if err != nil {
		// ゲストユーザーだった場合以外はエラー
		if err.ErrorCode != NotFoundUser {
			return true
		}
	}
	return false
}
