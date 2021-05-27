// +build wireinject

package wire

import (
	"go-vue-next-server/pkg/infrastructure/store"
	"go-vue-next-server/pkg/interface/handler"
	"go-vue-next-server/pkg/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

//InitializeUserAccountHandler userAccountの依存関係登録
func InitializeUserAccountHandler(db *gorm.DB) handler.UserAccountHandler {
	wire.Build(
		handler.NewUserAccountHandler,
		usecase.NewUserAccountUsecase,
		store.NewUserAccountStore,
	)
	return nil
}
