package controller

import (
	"app/model"
	"app/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// IUserController はUserControllerで実装されるインターフェースです。
type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUseCase
}

// NewUserController はUserControllerを初期化します。
func NewUserController(uu usecase.IUserUseCase) IUserController {
	return &userController{uu}
}

// SignUp はユーザのサインアップを行います。
func (uc *userController) SignUp(c echo.Context) error {
	// POSTされたJSONをmodel.User型にデコードする
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザを登録し、登録されたユーザ情報を返す
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// 登録されたユーザ情報を返す
	return c.JSON(http.StatusCreated, userRes)
}

// Login はユーザのログインを行います。
func (uc *userController) Login(c echo.Context) error {
	// POSTされたJSONをmodel.User型にデコードする
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザのログインを行い、トークンを返す
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Cookieにトークンを設定する
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// Logout はユーザのログアウトを行います。
func (uc *userController) Logout(c echo.Context) error {
	// Cookieからトークンを削除する
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// CsrfToken はCSRFトークンを返します。
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
