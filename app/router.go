package app

import (
	"database/sql"

	"github.com/adycahyoputro/manage-bank/delivery"
	"github.com/adycahyoputro/manage-bank/repository"
	"github.com/adycahyoputro/manage-bank/usecase"
	"github.com/adycahyoputro/manage-bank/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepository := repository.NewUserRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	entryRepository := repository.NewEntryRepository(db)
	transferRepository := repository.NewTransferRepository(db)
	transactionRepository := repository.NewTransactionRepository(db, userRepository, accountRepository, entryRepository, transferRepository)

	// userUsecase := usecase.NewUserUsecase(userRepository)
	accountUsecase := usecase.NewAccountUsecase(accountRepository)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository, userRepository, accountRepository)
	loginUsecase := usecase.NewLoginUsecase(userRepository, accountRepository)

	transactionDelivery := delivery.NewTransactionDelivery(transactionUsecase, accountUsecase)
	loginDelivery := delivery.NewLoginDelivery(loginUsecase)

	r.POST("/login", loginDelivery.Login)
	r.POST("/register", transactionDelivery.CreateUserAccount)

	group_router := r.Group("manage-bank")
	group_router.Use(utils.AuthMiddleware(utils.GetEnv("JWT_KEY")))
	{
		group_router.POST("/entry", transactionDelivery.CreateMainEntry)
		group_router.POST("/transfer", transactionDelivery.CreateMainTransfer)
		group_router.POST("/logout", loginDelivery.Logout)
	}

	return r
}
