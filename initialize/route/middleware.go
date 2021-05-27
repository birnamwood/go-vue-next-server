package route

import (
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// systemAuth 各handler処理前にHeaderのトークンの認証を行う
func systemAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		zap.S().Info("外部システムから送られたトークンのチェック開始")
		t := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		zap.S().Info(t)

		return next(c)
	}
}
