package utils

import (
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(ctx *gin.Context, err error) {
	switch err.Error() {
	case models.ErrorUserDuplicate.Error():
		ctx.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
	case models.ErrorUserWrongPassword.Error():
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case models.ErrorInternalServer.Error():
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	case models.ErrorUserNotFound.Error():
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

}
