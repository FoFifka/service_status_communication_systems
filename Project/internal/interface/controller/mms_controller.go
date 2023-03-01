package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type mmsController struct {
	MMSInteractor interactor.MMSInteractor
}

type MMSController interface {
	GetMMSData(ctx echo.Context) error
}

func NewMMSController(mi interactor.MMSInteractor) MMSController {
	return &mmsController{mi}
}

func (mc *mmsController) GetMMSData(ctx echo.Context) error {
	mmsResponse, err := mc.MMSInteractor.Get()
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusOK, mmsResponse)
}
