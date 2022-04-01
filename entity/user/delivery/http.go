package delivery

import (
	"be-golang-echo/entity/user"
	"be-golang-echo/entity/user/usecase"
	"be-golang-echo/utils"
	"be-golang-echo/utils/authentication"
	"be-golang-echo/utils/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHttpDelivery struct {
	UserUsecase usecase.UserUsecase
}

func NewHttpDelivery(e *echo.Echo, u usecase.UserUsecase) {
	handler := &UserHttpDelivery{
		UserUsecase: u,
	}

	api := e.Group("api/v1/user")
	api.POST("/register", handler.UserRegister)
	api.POST("/login", handler.UserLogin)
	api.GET("/profile", handler.GetProfile, authentication.IsAuthenticated())
	api.POST("/token", handler.RequestToken)

}

func (u *UserHttpDelivery) UserRegister(c echo.Context) error {
	resp := new(user.ResponseModel)
	userObj := new(user.UserModel)
	err := c.Bind(userObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	res, err := u.UserUsecase.UserRegister(userObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.REGISTER_SUCCESS
	resp.Data = res
	return c.JSON(http.StatusOK, resp)
}

func (u *UserHttpDelivery) UserLogin(c echo.Context) error {
	userObj := new(user.UserLoginModel)
	err := c.Bind(userObj)
	resp := new(user.ResponseModel)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	token, err := u.UserUsecase.UserLogin(userObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.LOGIN_SUCCESS
	resp.Data = token
	return c.JSON(http.StatusOK, resp)
}

func (u *UserHttpDelivery) GetProfile(c echo.Context) error {
	resp := new(user.ResponseModel)
	reqToken := c.Request().Header.Get(echo.HeaderAuthorization)
	id, err := jwt.GetIDfromToken(reqToken)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	art, err := u.UserUsecase.GetProfile(id)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.OK
	resp.Data = art
	return c.JSON(http.StatusOK, resp)
}

func (u *UserHttpDelivery) RequestToken(c echo.Context) error {
	resp := new(user.ResponseModel)
	userObj := new(user.TokenRequestModel)
	err := c.Bind(userObj)

	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	id, err := jwt.GetIDfromRefreshToken(userObj.RefreshToken)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	token, err := u.UserUsecase.RequestToken(id)

	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.OK
	resp.Data = token
	return c.JSON(http.StatusOK, resp)
}
