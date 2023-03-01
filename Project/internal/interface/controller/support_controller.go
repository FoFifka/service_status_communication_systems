package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type supportController struct {
	SupportInteractor interactor.SupportInteractor
}

type SupportController interface {
	GetSupportData(ctx echo.Context) error
}

func NewSupportController(si interactor.SupportInteractor) SupportController {
	return &supportController{si}
}

func (sc *supportController) GetSupportData(ctx echo.Context) error {
	sd, err := sc.SupportInteractor.Get()
	if err != nil {
		return err
	}
	return ctx.JSONBlob(http.StatusOK, sd)
}
