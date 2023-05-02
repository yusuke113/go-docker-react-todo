package usecase

import (
	"app/model"      // modelパッケージをインポート
	"app/repository" // repositoryパッケージをインポート
	"os"             // OS環境変数の読み取りに必要
	"time"           // 時間関連処理に必要

	"github.com/golang-jwt/jwt/v4" // JWT関連処理に必要
	"golang.org/x/crypto/bcrypt"   // bcrypt関連処理に必要
)

// IUserUseCaseはユーザ関連のユースケースで使用するメソッドのインターフェース
type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// userUseCaseはユーザ関連のユースケースの実装
type userUseCase struct {
	// IUserRepositoryを実装した構造体のインスタンス
	ur repository.IUserRepository
}

// NewUserUseCaseはuserUseCase構造体を生成するためのコンストラクタ
func NewUserUseCase(ur repository.IUserRepository) IUserUseCase {
	return &userUseCase{ur}
}

// SignUpはユーザ登録のユースケースの実装
func (uu *userUseCase) SignUp(user model.User) (model.UserResponse, error) {
	// ユーザパスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	// ハッシュ化したパスワードを格納したUser構造体を作成
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// UserRepositoryを利用して、新規ユーザをDBに登録
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// 登録後のユーザ情報をレスポンスとして返却
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

// Loginはログイン認証のユースケースの実装
func (uu *userUseCase) Login(user model.User) (string, error) {
	// ログインユーザーがデータベース上に存在するか確認するために、
	// メールアドレスをキーにユーザー情報を取得します。
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// 入力されたパスワードとデータベースから取得したハッシュ化パスワードを比較し、
	// 一致しなければエラーを返します。
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// ユーザー情報を元にJWTトークンを生成し、秘密鍵で署名した文字列として返します。
	// トークンのペイロードには、ユーザーIDとトークンの有効期限が含まれます。
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	// 環境変数から秘密鍵を取得し、トークンを文字列に変換します。
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	// トークン文字列を返します。
	return tokenString, nil
}
