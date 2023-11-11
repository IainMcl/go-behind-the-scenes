package api

import (
	"fmt"
	"net/http"

	"github.com/IainMcl/go-behind-the-scenes/internal/logging"
	"github.com/IainMcl/go-behind-the-scenes/internal/util"
	"github.com/IainMcl/go-behind-the-scenes/models"
	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func Login(c *gin.Context) {
	var data interface{}

	var u user
	if err := c.BindJSON(&u); err != nil {
		logging.Error("Error creating user object: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	if (u.Username == "" && u.Email == "") || u.Password == "" {
		logging.Debug("Username, Email or Password is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Username and Password or Email and Password are required",
		})
		return
	}

	if models.TryLogin(u.Username, u.Email, u.Password) {

		token, err := util.GenerateToken(u.Username)
		if err != nil {
			data = err.Error()
			logging.Error("Error generating token: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": "Error logging in please try again",
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

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logging.Error("Error binding JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		logging.Debug("Username, Password, or Email is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Username, Password and Email are required",
		})
		return
	}

	// Check if email already exists in the database
	if user, err := models.GetUserByEmail(user.Email); err == nil {
		logging.Warn(fmt.Sprintf("User with email %s already exists", user.Email))
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Email already exists",
		})
		return
	}

	var userModel models.User = models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	if err := models.AddUser(&userModel); err != nil {
		logging.Error("Error adding user: ", user.Username)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": "Error adding user",
		})
		return
	}

	logging.Info("User added: ", user.Username)
	msg := fmt.Sprintf("User %s added", user.Username)
	c.JSON(http.StatusOK, gin.H{
		"data": msg,
	})
}

func Logout(c *gin.Context) {
	data := "Logout"
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
