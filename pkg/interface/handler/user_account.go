package handler

import (
	"go-vue-next-server/pkg/domain/json/structure"
	"go-vue-next-server/pkg/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

//UserAccountHandler interface
type UserAccountHandler interface {
	CheckUserAccount(echo.Context) error
	GetTest(echo.Context) error
}

//userAccountHandler struct
type userAccountHandler struct {
	userAccountUsecase usecase.UserAccountUsecase
}

//NewUserAccountHandler New
func NewUserAccountHandler(
	userAccountUsecase usecase.UserAccountUsecase,
) UserAccountHandler {
	return &userAccountHandler{
		userAccountUsecase: userAccountUsecase,
	}
}

// GetTest 疎通確認
func (h *userAccountHandler) GetTest(c echo.Context) error {
	v := map[string]string{"Result": "OK"}
	return c.JSON(http.StatusOK, v)
}

// CheckUserAccount function
func (h *userAccountHandler) CheckUserAccount(c echo.Context) error {
	result := new(structure.CheckAccountResult)

	result.NewUserAccount = false
	return c.JSON(http.StatusOK, result)
}
