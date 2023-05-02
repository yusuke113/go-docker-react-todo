package repository

import (
	"app/model"

	"gorm.io/gorm"
)

// IUserRepositoryはユーザーデータを永続化するためのリポジトリのインターフェースです
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// userRepositoryはユーザーデータを永続化するためのリポジトリの実装です
type userRepository struct {
	db *gorm.DB
}

// NewUserRepositoryはuserRepositoryのインスタンスを生成します
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// GetUserByEmailは指定されたemailアドレスを持つユーザーデータを取得します
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUserは新しいユーザーデータを作成します
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
