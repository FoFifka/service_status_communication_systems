package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type resultController struct {
	ResultInteractor interactor.ResultInteractor
}

type ResultController interface {
	GetResultData(ctx echo.Context) error
}

func NewResultController(ri interactor.ResultInteractor) ResultController {
	return &resultController{ri}
}

func (rc *resultController) GetResultData(ctx echo.Context) error {
	r, err := rc.ResultInteractor.Get()
	if err != nil {
		return ctx.JSONBlob(http.StatusServiceUnavailable, r)
	}
	return ctx.JSONBlob(http.StatusOK, r)
}
