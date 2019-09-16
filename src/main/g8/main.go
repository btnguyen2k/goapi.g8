/*
Application Server bootstrapper.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@since template-v0.1.0
*/
package main

import (
	"main/src/goems"
	"main/src/samples"
	"main/src/samples_crud_mongodb"
	"main/src/samples_crud_pgsql"
	"math/rand"
	"time"
)

func main() {
	// it is a good idea to initialize random seed
	rand.Seed(time.Now().UnixNano())

	// start Echo server with custom bootstrapper
	// bootstrapper routine is passed the echo.Echo instance as argument, and also has access to
	// - Application configurations via global variable goems.AppConfig
	// - itineris.ApiRouter instance via global variable goems.ApiRouter
	goems.Start(samples.Bootstrapper, samples_crud_mongodb.Bootstrapper, samples_crud_pgsql.Bootstrapper)
}