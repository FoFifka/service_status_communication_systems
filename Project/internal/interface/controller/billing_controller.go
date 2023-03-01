package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type billingController struct {
	BillingInteractor interactor.BillingInteractor
}

type BillingController interface {
	GetBillingData(ctx echo.Context) error
}

func NewBillingController(bi interactor.BillingInteractor) BillingController {
	return &billingController{bi}
}

func (bc *billingController) GetBillingData(ctx echo.Context) error {
	billingResponse, err := bc.BillingInteractor.Get()
	if err != nil {
		return err
	}
	return ctx.JSONBlob(http.StatusOK, billingResponse)
}
