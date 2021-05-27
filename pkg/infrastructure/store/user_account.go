package store

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"go-vue-next-server/initialize/alert"
	"go-vue-next-server/pkg/domain/model"
	"go-vue-next-server/pkg/domain/repository"
	"go-vue-next-server/pkg/general/myerr"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//UserAccountStore struct
type UserAccountStore struct {
	db *gorm.DB
}

//NewUserAccountStore 初期化
func NewUserAccountStore(database *gorm.DB) repository.UserAccountRepository {
	return &UserAccountStore{db: database}
}

//Create UserAccountCreate
func (p *UserAccountStore) Create(userAccount *model.UserAccount) (*model.UserAccount, error) {
	if err := p.db.Create(&userAccount).Error; err != nil {
		key, _ := json.Marshal(&userAccount)
		alert.SendMail(errors.Wrap(err, "UserAccount情報の作成に失敗しました。"), string(debug.Stack()), string(key))
		return nil, myerr.ErrCreateUserAccountFailed
	}
	return userAccount, nil
}

// Delete ユーザーの削除
func (p *UserAccountStore) Delete(id int) error {
	if err := p.db.Where("id = ?", id).Delete(&model.UserAccount{}).Error; err != nil {
		key, _ := json.Marshal(map[string]int{"UserAccountID": id})
		alert.SendMail(errors.Wrap(err, "UserAccountの削除に失敗しました。"), string(debug.Stack()), string(key))
		return myerr.ErrDeleteUserAccountFailed
	}
	return nil
}

//FindByID IDによるUserAccount検索
func (p *UserAccountStore) FindByID(id int) (*model.UserAccount, error) {
	userAccount := &model.UserAccount{}
	if err := p.db.Where("ID = ?", id).First(&userAccount).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			key, _ := json.Marshal(map[string]string{"UserAccountID": fmt.Sprint(id)})
			alert.SendMail(errors.Wrap(err, "IDでのUserAccount検索に失敗しました。"), string(debug.Stack()), string(key))
			return nil, myerr.ErrSearchUserAccountFailed
		}
		return nil, myerr.ErrUserAccountNotFound
	}
	return userAccount, nil
}
