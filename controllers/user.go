package controllers

import (
	"fmt"
	"net/http"

	"github.com/bhoopendrau/tailscale-ui-backend/forms"
	"github.com/bhoopendrau/tailscale-ui-backend/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.GetByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u UserController) SignUp(c *gin.Context) {
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	var useForm forms.UserSignup
	if err := c.BindJSON(&useForm); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body!", "error": err})
		return
	}
	if _, err := userModel.Signup(useForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to create document!", "error": err})
		return
	}

}
