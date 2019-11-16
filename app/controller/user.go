package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// If `GET`, only `Form` binding engine (`query`) used.
// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
type RegisterParam struct {
	Email    string `form:"email" binding:"required,email"`
	Username string `form:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginParam struct {
	Type     string `form:"type" binding:"required,oneof=email username"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func HttpHandlerLogin(c *gin.Context) {
	param := LoginParam{}
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	c.JSON(200, gin.H{"status": "you are logged in"})

	// token, err := managers.AccountLogin(account.Email, account.Password)
	// if err != nil {
	// 	c.JSON(http.StatusOK, base.Fail(err.Error()))
	// 	return
	// }
	// cookie := &http.Cookie{
	// 	Name:     "token",
	// 	Value:    base64.StdEncoding.EncodeToString([]byte(token)),
	// 	Path:     "/",
	// 	HttpOnly: true,
	// }
	//
	// http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, "ok")
}

//
// func httpHandlerRegister(c *gin.Context) {
// 	account := AccountParam{}
// 	err := c.Bind(&account)
// 	if err != nil {
// 		panic(err)
// 	}
// 	userId, err := managers.AccountRegister(account.Email, account.Password)
// 	if err != nil {
// 		c.JSON(http.StatusOK, base.Fail(err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, base.Success(userId))
// }