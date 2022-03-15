package delivery

import (
	"be-golang-echo/entity/pengerjaan"
	"be-golang-echo/entity/pengerjaan/usecase"
	"be-golang-echo/utils"
	"be-golang-echo/utils/config_variable"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type PengerjaanHttpDelivery struct {
	PengerjaanUsecase usecase.PengerjaanUsecase
}

func NewHttpDelivery(e *echo.Echo, p usecase.PengerjaanUsecase) {
	handler := &PengerjaanHttpDelivery{
		PengerjaanUsecase: p,
	}

	api := e.Group("api/v1/pengerjaans")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config_variable.Secret),
	}))
	api.GET("/:id", handler.GetSingle)
	api.PUT("/:id", handler.Update)
	api.DELETE("/:id", handler.Delete)
	api.GET("", handler.GetList)
	api.POST("", handler.Add)
}

func (p *PengerjaanHttpDelivery) GetList(c echo.Context) error {
	resp := new(pengerjaan.ResponseModel)
	// reqToken := c.Request().Header.Get(echo.HeaderAuthorization)

	// admin, err := jwt.VerifyAdminToken(reqToken)
	// if !admin {
	// 	resp.Status = utils.FAILED
	// 	resp.Message = err.Error()
	// 	return c.JSON(http.StatusInternalServerError, resp)
	// }

	// id, err := jwt.GetIDfromToken(reqToken)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, common.InternalServerErrorResponse(err.Error()))
	// }

	filters := new(pengerjaan.Filter)

	if len(c.QueryParam("limit")) == 0 {
		resp.Status = utils.FAILED
		resp.Message = utils.LIMIT_EMPTY
		return c.JSON(http.StatusInternalServerError, resp)
	}

	filters.Limit = c.QueryParam("limit")

	art, err := p.PengerjaanUsecase.GetList(filters)
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

func (p *PengerjaanHttpDelivery) Add(c echo.Context) error {
	resp := new(pengerjaan.ResponseModel)

	pengerjaanObj := new(pengerjaan.PengerjaanModel)
	err := c.Bind(pengerjaanObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}

	_, err = p.PengerjaanUsecase.Add(pengerjaanObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.ADD_NEW_SUCCESS
	return c.JSON(http.StatusOK, resp)
}

func (p *PengerjaanHttpDelivery) Update(c echo.Context) error {
	resp := new(pengerjaan.ResponseModel)

	idP, err := strconv.Atoi(c.Param("id"))
	id := int(idP)
	pengerjaanObj := new(pengerjaan.PengerjaanModel)
	err = c.Bind(pengerjaanObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	_, err = p.PengerjaanUsecase.Update(id, pengerjaanObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.UPDATE_SUCCESS
	return c.JSON(http.StatusOK, resp)
}

func (p *PengerjaanHttpDelivery) Delete(c echo.Context) error {
	resp := new(pengerjaan.ResponseModel)

	idP, err := strconv.Atoi(c.Param("id"))
	id := int(idP)

	_, err = p.PengerjaanUsecase.Delete(id)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.DELETE_SUCCESS
	return c.JSON(http.StatusOK, resp)
}

func (p *PengerjaanHttpDelivery) GetSingle(c echo.Context) error {
	resp := new(pengerjaan.ResponseModel)

	idP, err := strconv.Atoi(c.Param("id"))
	id := int(idP)

	data, err := p.PengerjaanUsecase.GetSingle(id)

	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.OK
	resp.Data = data
	return c.JSON(http.StatusOK, resp)
}
