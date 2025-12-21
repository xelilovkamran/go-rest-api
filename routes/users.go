package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)


func signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user data."})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}


func login(c *gin.Context) {
	var credentials struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		validationErrors := utils.FormatValidationError(err)
		if len(validationErrors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse credentials."})
		return
	}


	user, err := models.GetUserByEmail(credentials.Email)
	if err != nil || !utils.VerifyPassword(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password."})
		return
	}

	jwtToken, err := utils.CreateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create JWT token.", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}