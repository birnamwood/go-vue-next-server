package repository

import (
	"go-vue-next-server/pkg/domain/model"
)

//UserAccountRepository text
type UserAccountRepository interface {
	Create(userAccount *model.UserAccount) (*model.UserAccount, error)
	Delete(id int) error
	FindByID(id int) (*model.UserAccount, error)
}
