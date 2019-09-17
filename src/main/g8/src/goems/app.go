package goems

import (
	"encoding/json"
	hocon "github.com/go-akka/configuration"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"main/src/itineris"
	"main/src/utils"
	"os"
	"strconv"
	"time"
)

const (
	defaultConfigFile = "./config/application.conf"
)

var (
	AppConfig *hocon.Config
	ApiRouter *itineris.ApiRouter
)

/*
Start bootstraps the application.
*/
func Start(bootstrappers ...IBootstrapper) {
	var err error

	// load application configurations
	AppConfig = initAppConfig()
	httpHeaderAppId = AppConfig.GetString("api.http.header_app_id")
	httpHeaderAccessToken = AppConfig.GetString("api.http.header_access_token")

	// setup api-router
	ApiRouter = itineris.NewApiRouter()

	// initialize "Location"
	utils.Location, err = time.LoadLocation(AppConfig.GetString("timezone"))
	if err != nil {
		panic(err)
	}

	// bootstrapping
	if bootstrappers != nil {
		for _, b := range bootstrappers {
			log.Println("Bootstrapping", b)
			err := b.Bootstrap()
			if err != nil {
				log.Println(err)
			}
		}
	}

	// initialize and start echo server
	e := initEcho()
	listenAddr := AppConfig.GetString("api.http.listen_addr", "127.0.0.1")
	listenPort := AppConfig.GetInt32("api.http.listen_port", 8080)
	log.Println("Starting " + AppConfig.GetString("app.name") + " v" + AppConfig.GetString("app.version") + "...")
	e.Logger.Fatal(e.Start(listenAddr + ":" + strconv.Itoa(int(listenPort))))
}

func initAppConfig() *hocon.Config {
	configFile := os.Getenv("APP_CONFIG")
	if configFile == "" {
		log.Printf("No environment APP_CONFIG found, fallback to [%s]", defaultConfigFile)
		configFile = defaultConfigFile
	}
	return loadAppConfig(configFile)
}

func initEcho() *echo.Echo {
	e := echo.New()
	requestTimeout := AppConfig.GetTimeDuration("api.request_timeout", time.Duration(0))
	if requestTimeout > 0 {
		e.Server.ReadTimeout = requestTimeout
	}
	bodyLimit := AppConfig.GetByteSize("api.max_request_size")
	if bodyLimit != nil && bodyLimit.Int64() > 0 {
		e.Use(middleware.BodyLimit(bodyLimit.String()))
	}

	// register API http endpoints
	hasEndpoints := false
	confV := AppConfig.GetValue("api.http.endpoints")
	if confV != nil && confV.IsObject() {
		for uri, uriO := range confV.GetObject().Items() {
			if uriO.IsObject() && !uriO.IsEmpty() {
				hasEndpoints = true
				e.Any(uri, apiHttpHandler)
				for httpMethod, apiName := range uriO.GetObject().Items() {
					registerHttpHandler(uri, httpMethod, apiName.GetString())
				}
			}
		}
	}
	js, _ := json.Marshal(httpRoutingMap)
	log.Println("API http endpoints: " + string(js))
	if !hasEndpoints {
		log.Println("No valid HTTP endpoints defined at key [api.http.endpoints].")
	}
	return e
}
