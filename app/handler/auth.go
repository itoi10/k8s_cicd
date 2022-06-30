/*
ユーザー登録、ユーザー認証処理
*/
package handler

import (
	"echo_sample/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type jwtCustomClaims struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func Signup(c echo.Context) error {
	// パラメータの値を構造体に格納
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// パラメータチェック
	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	// ユーザー名重複チェック
	if u := model.FindUser(&model.User{Name: user.Name}); u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	// 新規ユーザー作成
	model.CreateUser(user)

	// パスワードを空文字にしたユーザー情報返却
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	// パラメータの値を構造体に格納
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// ユーザーが存在しないかパスワードが違う
	user := model.FindUser(&model.User{Name: u.Name})
	if user.ID == 0 || user.Password != u.Password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}

	// クレーム設定
	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// JWTトークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	// トークン返却
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
