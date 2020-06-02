package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func PostRegister(c *gin.Context) {
	var userinfo auth
	if err := c.BindJSON(&userinfo) err != nil {
		return
	}
	valid := validation.Validation{}
	valid.Required(userinfo.username, "username").Message("用户名不能为空")
	valid.Required(userinfo.password, "password").Message("密码不能为空")

	// code := e.INVALID_PARAMS
	// data := make(map[string]interface{})
	// if !valid.HasErrors() {
	// 	if !models.ExistAuthByUsername(username) {
	// 		code = e.SUCCESS
	// 		models.AddAuth(username, password, createdBy)
	// 		token, err := util.GenerateToken(username, password)
	// 		if err != nil {
	// 			code = e.ERROR_AUTH_TOKEN
	// 		} else {
	// 			data["token"] = token
	// 		}
	// 	} else {
	// 		code = e.ERROR_EXIST_USERNAME
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{
		// "code": code,
		// "msg":  e.GetMsg(code),
		"data": userinfo,
	})
}
