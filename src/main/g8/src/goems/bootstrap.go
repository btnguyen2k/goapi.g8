package goems

/*
IBootstrapper defines an interface for application to hook bootstrapping routines.

Bootstrapper has access to:
- Application configurations via global variable goems.AppConfig
- itineris.ApiRouter instance via global variable goems.ApiRouter
*/
type IBootstrapper interface {
    Bootstrap() error
}
