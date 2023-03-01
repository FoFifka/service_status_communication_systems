package main

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"skillsbox-diploma/internal/infrastructure/router"
	"skillsbox-diploma/internal/registry"
	"skillsbox-diploma/pkg/app_settings"
)

func main() {
	app_settings.AppSettings = app_settings.SetSettings()

	host := app_settings.AppSettings.Configuration.Host
	port := app_settings.AppSettings.Configuration.Port

	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("serv started...")
	if err := e.Start(host + ":" + port); err != nil {
		log.Fatalln(err)
	}
}
