package server

import (
	"github.com/avtara/boilerplate-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

func (so *svObject) handlerGetLastLogin(ctx *gin.Context) {
	var json models.GetLastLoginRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	result, err := so.UserUsecase.GetLastLogin(ctx.Request.Context(), json)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	return
}

func (so *svObject) handlerRegister(ctx *gin.Context) {
	var json models.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	result, err := so.UserUsecase.Register(ctx.Request.Context(), json)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
		return
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	return
}
