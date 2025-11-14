package routes

import (
	"net/http"

	"example.com/RestAPI/models"
	"example.com/RestAPI/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "event": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user.", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful.", "token": token})
}
