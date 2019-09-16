package samples_api_filters

import (
	"main/src/goems"
	"main/src/itineris"
	"os"
)

type MyBootstrapper struct {
	name string
}

var Bootstrapper = &MyBootstrapper{name: "samples_api_filters"}

func (b *MyBootstrapper) Bootstrap() error {
	var apiFilter itineris.IApiFilter = nil
	appName := goems.AppConfig.GetString("app.name")
	appVersion := goems.AppConfig.GetString("app.version")

	apiFilter = itineris.NewAddPerfInfoFilter(goems.ApiRouter, apiFilter)
	apiFilter = itineris.NewLoggingFilter(goems.ApiRouter, apiFilter, itineris.NewWriterPerfLogger(os.Stderr, appName, appVersion))
	apiFilter = itineris.NewLoggingFilter(goems.ApiRouter, apiFilter, itineris.NewWriterRequestLogger(os.Stdout, appName, appVersion))

	goems.ApiRouter.SetApiFilter(apiFilter)
	return nil
}
