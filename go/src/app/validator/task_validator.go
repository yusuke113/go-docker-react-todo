package validator

import (
	"app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ITaskValidatorはタスクのバリデーションを行うためのインターフェースです
type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

// taskValidatorはタスクのバリデーションを行うための構造体です
type taskValidator struct{}

// NewTaskValidatorは新しいTaskValidatorのインスタンスを作成して返します
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// TaskValidateはタスクのバリデーションを行います
func (tv *taskValidator) TaskValidate(task model.Task) error {
	// タスクのバリデーションルールを定義して適用します
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("Title is required"),                  // タイトルは必須項目であることを検証します
			validation.RuneLength(1, 10).Error("limited max 10 characters"), // タイトルの文字数が1以上10以下であることを検証します
		),
	)
}
