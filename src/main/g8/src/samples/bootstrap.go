/*
Package samples provides samples API implementation.
*/
package samples

import (
    "main/src/goems"
    "main/src/itineris"
    "runtime"
    "strconv"
)

type MyBootstrapper struct {
    name string
}

var Bootstrapper = &MyBootstrapper{name: "samples"}

/*
Bootstrap implements goems.IBootstrapper.Bootstrap

Bootstrapper usually does:
- register api-handlers with the global ApiRouter
- other initializing work (e.g. creating DAO, initializing database, etc)
*/
func (b *MyBootstrapper) Bootstrap() error {
    initApiHandlers(goems.ApiRouter)
    return nil
}

/*
Setup API handlers: application register its api-handlers by calling router.SetHandler(apiName, apiHandlerFunc)

    - api-handler function must has the following signature: func (itineris.ApiContext, itineris.ApiAuth, itineris.ApiParams) *itineris.ApiResult
*/
func initApiHandlers(router *itineris.ApiRouter) {
    router.SetHandler("info", apiInfo)
    router.SetHandler("echo", apiEcho)
}

/*
API handler "info"
*/
func apiInfo(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    appInfo := map[string]interface{}{
        "name":        goems.AppConfig.GetString("app.name"),
        "shortname":   goems.AppConfig.GetString("app.shortname"),
        "version":     goems.AppConfig.GetString("app.version"),
        "description": goems.AppConfig.GetString("app.desc"),
    }
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    result := map[string]interface{}{
        "app": appInfo,
        "memory": map[string]interface{}{
            "alloc":     m.Alloc,
            "alloc_str": strconv.FormatFloat(float64(m.Alloc)/1024.0/1024.0, 'f', 1, 64) + " MiB",
            "sys":       m.Sys,
            "sys_str":   strconv.FormatFloat(float64(m.Sys)/1024.0/1024.0, 'f', 1, 64) + " MiB",
            "gc":        m.NumGC,
        },
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}

/*
API handler "echo"
*/
func apiEcho(ctx *itineris.ApiContext, auth *itineris.ApiAuth, params *itineris.ApiParams) *itineris.ApiResult {
    result := map[string]interface{}{
        "context": ctx.GetAllContextValues(),
        "auth": map[string]interface{}{
            "app_id":       auth.GetAppId(),
            "access_token": auth.GetAccessToken(),
        },
        "params": params.GetAllParams(),
    }
    return itineris.NewApiResult(itineris.StatusOk).SetData(result)
}
