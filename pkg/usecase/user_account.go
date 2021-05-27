package usecase

import (
	"go-vue-next-server/pkg/domain/model"
	"go-vue-next-server/pkg/domain/repository"
)

//UserAccountUsecase interface
type UserAccountUsecase interface {
	FindByID(id int) (*model.UserAccount, error)
	CreateUserAccount(userAccount *model.UserAccount) (*model.UserAccount, error)
}

//userAccountUsecase struct
type userAccountUsecase struct {
	userAccountRepository repository.UserAccountRepository
}

//NewUserAccountUsecase New
func NewUserAccountUsecase(
	userAccountRepository repository.UserAccountRepository,
) UserAccountUsecase {
	return &userAccountUsecase{
		userAccountRepository: userAccountRepository,
	}
}

// FindByID IDでのユーザー検索
func (u *userAccountUsecase) FindByID(id int) (*model.UserAccount, error) {
	user, err := u.userAccountRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUserAccount ユーザーの新規登録
func (u *userAccountUsecase) CreateUserAccount(user *model.UserAccount) (*model.UserAccount, error) {
	user, err := u.userAccountRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
