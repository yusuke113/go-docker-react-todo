package validator

import (
	"app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// IUserValidatorはユーザーのバリデーションを行うためのインターフェースです
type IUserValidator interface {
	UserValidate(user model.User) error
}

// userValidatorはユーザーのバリデーションを行うための構造体です
type userValidator struct{}

// NewUserValidatorは新しいUserValidatorのインスタンスを作成して返します
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// UserValidateはユーザーのバリデーションを行います
func (uv *userValidator) UserValidate(user model.User) error {
	// ユーザーのバリデーションルールを定義して適用します
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required"),                  // メールアドレスは必須項目であることを検証します
			validation.RuneLength(1, 30).Error("limited max 30 characters"), // メールアドレスの文字数が1以上30以下であることを検証します
			is.Email.Error("is not valid email format"),                     // メールアドレスのフォーマットが正しいことを検証します
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required"),                     // パスワードは必須項目であることを検証します
			validation.RuneLength(6, 30).Error("limited min 6 max 30 characters"), // パスワードの文字数が6以上30以下であることを検証します
		),
	)
}
