package delivery

import (
	"net/http"

	"github.com/adycahyoputro/manage-bank/model/dto"
	"github.com/adycahyoputro/manage-bank/usecase"
	"github.com/adycahyoputro/manage-bank/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TransactionDelivery interface {
	CreateUserAccount(*gin.Context)
	CreateMainTransfer(*gin.Context)
	CreateMainEntry(*gin.Context)
}

type transactionDelivery struct {
	transactionUsecase usecase.TransactionUsecase
	accountUsecase     usecase.AccountUsecase
}

func NewTransactionDelivery(
	transactionUsecase usecase.TransactionUsecase,
	accountUsecase usecase.AccountUsecase) TransactionDelivery {
	return &transactionDelivery{
		transactionUsecase: transactionUsecase,
		accountUsecase:     accountUsecase}
}

func (delivery *transactionDelivery) CreateUserAccount(ctx *gin.Context) {
	var userRequest dto.UserRequest

	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		result := dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROsR",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	newUserAccount, err := delivery.transactionUsecase.CreateUserAccount(&userRequest)
	if err != nil {
		result := utils.CheckError(userRequest.Email, err)
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	result := dto.Response{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   newUserAccount,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (delivery *transactionDelivery) CreateMainTransfer(ctx *gin.Context) {
	var transferRequest dto.TransferRequest

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userId := claims["id"].(string)

	err := ctx.ShouldBindJSON(&transferRequest)
	if err != nil {
		result := dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROsR",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	getAccountByUserID, _ := delivery.accountUsecase.FindAccountByUserID(userId)
	fromAccountID := getAccountByUserID.ID

	newUserAccount, err := delivery.transactionUsecase.CreateMainTransfer(&transferRequest, fromAccountID)
	if err != nil {
		result := utils.CheckError(transferRequest.ToAccountID, err)
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	result := dto.Response{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   newUserAccount,
	}
	ctx.JSON(http.StatusCreated, result)
}

func (delivery *transactionDelivery) CreateMainEntry(ctx *gin.Context) {
	var entryRequest dto.EntryRequest

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	userId := claims["id"].(string)

	err := ctx.ShouldBindJSON(&entryRequest)
	if err != nil {
		result := dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	getAccountByUserID, _ := delivery.accountUsecase.FindAccountByUserID(userId)
	accountID := getAccountByUserID.ID
	balance := getAccountByUserID.Balance

	newEntry, err := delivery.transactionUsecase.CreateMainEntry(&entryRequest, accountID, balance)
	if err != nil {
		result := utils.CheckError(accountID, err)
		ctx.JSON(http.StatusInternalServerError, result)
		return
	}

	result := dto.Response{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   newEntry,
	}
	ctx.JSON(http.StatusCreated, result)
}
