package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLogin post login request
func GetLogin (c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"title": "Login",
	})
}

// PostLogin post login request
func PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	//TODO: database から取得する
	var users = gin.H{
		"hoge" : gin.H{"password": "huga"},
	}
	login := "ng"
	pass := ""
	if user, ok := users[email]; ok {
		pass = user.(gin.H)["password"].(string)
		if pass == password {
			login = "ok"
		}
	}

//	c.JSON(http.StatusOK, gin.H{
//		"login": login,
//		"email": email,
//		"password": password,
//		"password2": pass,
//	})

	c.HTML(http.StatusOK, "login", gin.H{
		"title": "Login",
	})
}
