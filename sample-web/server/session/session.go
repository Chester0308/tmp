package session

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sample-web/server/model"
)

//type SessionData struct {
//	User *User
//}

func GetUser(c *gin.Context) model.User {
	session := sessions.Default(c)
	id := session.Get("userID")
	name := session.Get("userName")
}

func SetUser(c *gin.Context, user model.User) {
	session := sessions.Default(c)
	session.Set("userID", user.UserID)
	session.Set("userName", user.Name)
	session.Save()
}