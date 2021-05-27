package myerr

//ErrorCode 独自エラー型を定義
type ErrorCode struct {
	Code    int
	Message string
}

/// Code 999 = 外部システムからのエラー
var (
	// ErrInvalidValue 値がおかしい場合
	ErrInvalidValue = ErrorCode{Code: 1, Message: "値が不正です"}
	// ErrKeyValueRequired DB検索用のキー情報が空の場合
	ErrKeyValueRequired = ErrorCode{Code: 2, Message: "検索キーの値がありません"}
	// ErrJSONReceiveFailed Json受け取り失敗
	ErrJSONReceiveFailed = ErrorCode{Code: 3, Message: "送信されたJSONの受け取りに失敗しました"}
	// ErrUnauthorized Tokenの認証に失敗した場合
	ErrUnauthorized = ErrorCode{Code: 4, Message: "トークンの認証に失敗しました。"}
	// ErrCloseFailed ファイルクローズに失敗
	ErrCloseFailed = ErrorCode{Code: 5, Message: "ファイルが閉じられませんでした"}
	// ErrPathParameterRequired URLのパラメータが空の場合
	ErrPathParameterRequired = ErrorCode{Code: 6, Message: "URLのパラメータは必須です"}
	// ErrParseJSONFailed Jsonの解析、変換に失敗した場合
	ErrParseJSONFailed = ErrorCode{Code: 7, Message: "情報の受け取りに失敗しました"}

	// ErrCreateUserAccountFailed ユーザーアカウントの登録に失敗した場合
	ErrCreateUserAccountFailed = ErrorCode{Code: 103, Message: "ユーザーアカウントの登録に失敗しました"}
	// ErrDeleteUserAccountFailed ユーザーアカウントの検索に失敗した場合
	ErrDeleteUserAccountFailed = ErrorCode{Code: 305, Message: "ユーザーアカウントの削除に失敗しました"}

	// ErrUserAccountNotFound ユーザーが見つからない場合
	ErrUserAccountNotFound = ErrorCode{Code: 401, Message: "ユーザーアカウントが見つかりませんでした"}
	// ErrSearchUserAccountFailed UserAccountの検索に失敗した場合
	ErrSearchUserAccountFailed = ErrorCode{Code: 504, Message: "ユーザーアカウントの検索に失敗しました"}
)

// Error error型のインターフェースを満たすため実装
func (e ErrorCode) Error() string {
	return e.Message
}

// GetCode 数字のコードを返す
func (e ErrorCode) GetCode() int {
	return e.Code
}
