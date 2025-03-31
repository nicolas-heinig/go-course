package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		errorResponse("Could not parse user", http.StatusBadRequest, ctx, err)
		return
	}

	err = user.Save()

	if err != nil {
		errorResponse("Could not create user", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		errorResponse("Could not parse user", http.StatusBadRequest, ctx, err)
		return
	}

	err = user.ValidateCreds()

	if err != nil {
		errorResponse("Unauthorized!", http.StatusUnauthorized, ctx, err)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		errorResponse("Could not generate token", http.StatusInternalServerError, ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
