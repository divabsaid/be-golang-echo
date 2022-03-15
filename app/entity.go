package app

import (
	pengerjaanDelivery "be-golang-echo/entity/pengerjaan/delivery"
	pengerjaanRepo "be-golang-echo/entity/pengerjaan/repository"
	pengerjaanUsecase "be-golang-echo/entity/pengerjaan/usecase"

	userDelivery "be-golang-echo/entity/user/delivery"
	userRepo "be-golang-echo/entity/user/repository"
	userUsecase "be-golang-echo/entity/user/usecase"

	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitEntity(e *echo.Echo, dbConn *sql.DB) {
	pengerjaanRepo := pengerjaanRepo.NewMySQLPengerjaanRepository(dbConn)
	pengerjaanUsecase := pengerjaanUsecase.NewPengerjaanUseCase(pengerjaanRepo)
	pengerjaanDelivery.NewHttpDelivery(e, pengerjaanUsecase)

	userRepo := userRepo.NewMySQLUserRepository(dbConn)
	userUsecase := userUsecase.NewUserUseCase(userRepo)
	userDelivery.NewHttpDelivery(e, userUsecase)
}
