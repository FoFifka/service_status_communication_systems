package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/internal/usecase/interactor"
)

type smsController struct {
	SMSInteractor interactor.SMSInteractor
}

type SMSController interface {
	GetSMSData(ctx echo.Context) error
}

func NewSMSController(si interactor.SMSInteractor) SMSController {
	return &smsController{si}
}

func (sc *smsController) GetSMSData(ctx echo.Context) error {
	var sms *[]model.SMSData

	sms, err := sc.SMSInteractor.Get(sms)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, sms)
}
