package delivery

import (
	"net/http"

	"github.com/adycahyoputro/manage-bank/model/dto"
	"github.com/adycahyoputro/manage-bank/usecase"
	"github.com/adycahyoputro/manage-bank/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginDelivery interface {
	Login(*gin.Context)
	Logout(*gin.Context)
}

type loginDelivery struct {
	loginUsecase usecase.LoginUsecase
}

func NewLoginDelivery(
	loginUsecase usecase.LoginUsecase) LoginDelivery {
	return &loginDelivery{
		loginUsecase: loginUsecase,
	}
}

func (delivery *loginDelivery) Login(ctx *gin.Context) {
	var login dto.LoginRequest

	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		result := dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	tokenString, err := delivery.loginUsecase.Login(&login)
	if err != nil {
		result := utils.CheckError(login.Email, err)
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := dto.Response{
		Code:   http.StatusAccepted,
		Status: "ACCEPTED",
		Data:   tokenString,
	}
	ctx.JSON(http.StatusAccepted, result)
}

func (delivery *loginDelivery) Logout(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userId := claims["id"].(string)

	err := delivery.loginUsecase.Logout(userId)
	if err != nil {
		result := dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}
	result := dto.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, result)
}
