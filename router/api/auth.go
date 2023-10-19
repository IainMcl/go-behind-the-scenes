package api

import (
	"net/http"

	"github.com/IainMcl/go-behind-the-scenes/internal/util"
	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data interface{}

	var u user
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"data": err.Error(),
		})
	}

	if u.Username == "guest" && u.Password == "guest" {

		token, err := util.GenerateToken(u.Username, u.Password)
		if err != nil {
			data = err.Error()
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": data,
			})
			return
		}

		data = u.Username
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"data":  data,
		})
		return
	} else {
		data = "Unauthorized"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": data,
		})
		return
	}
}

func Logout(c *gin.Context) {
	data := "Logout"
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
