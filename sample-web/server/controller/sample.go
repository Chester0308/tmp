package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetIndex show Hello world !!
func GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "Hello world !!")
}

// GetFullName get request sample
func GetFullName(c *gin.Context) {
	fname := c.DefaultQuery("firstname", "Guest")
	lname := c.DefaultQuery("lastname", "Last")
	//lname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") と同じ
	c.String(http.StatusOK, "Hello %s %s !!", fname, lname)
}

// PostMessage post request sample
func PostMessage(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "Guest")

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"name":    name,
	})
}

// SetCookie cookie sample
func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("sample")
	if err != nil {
		cookie = "none"
		c.SetCookie("sample", "cookieValue", 3600, "/sample/set-cookie", "localhost", false, true)
	}

	c.JSON(http.StatusOK, gin.H{
		"value": cookie,
	})
}

// BasicAuth Basic Auth sample
func BasicAuth(c *gin.Context) {
	var admins = gin.H{
		"admin": gin.H{"email": "admin@example.com"},
		"hoge":  gin.H{"email": "hoge@huga.com"},
	}
	// BasicAuth ミドルウェアによって設定される
	user := c.MustGet(gin.AuthUserKey).(string)
	if admin, ok := admins[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "admin": admin})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "admin": "No admin data :("})
	}
}

// Html sample
func Html(c *gin.Context) {
	var admins = gin.H{
		"admin": gin.H{"email": "admin@example.com"},
		"hoge":  gin.H{"email": "hoge@huga.com"},
	}
	// BasicAuth ミドルウェアによって設定される
	user := c.MustGet(gin.AuthUserKey).(string)
	if admin, ok := admins[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "admin": admin})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "admin": "No admin data :("})
	}
}
