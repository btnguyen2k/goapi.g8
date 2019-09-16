package itineris

/*
ApiRouter is responsible to routing API call to handler.
*/
type ApiRouter struct {
    handlersMap map[string]func(*ApiContext, *ApiAuth, *ApiParams) *ApiResult
}

/*
NewApiRouter creates a new ApiRouter instance.
*/
func NewApiRouter() *ApiRouter {
    return &ApiRouter{handlersMap: map[string]func(*ApiContext, *ApiAuth, *ApiParams) *ApiResult{}}
}

/*
GetHandler returns an api-handler by name.
*/
func (router *ApiRouter) GetHandler(apiName string) func(*ApiContext, *ApiAuth, *ApiParams) *ApiResult {
    f, ok := router.handlersMap[apiName]
    if ok {
        return f
    }
    return nil
}

/*
SetHandler maps an handler to api name.
*/
func (router *ApiRouter) SetHandler(apiName string, handler func(*ApiContext, *ApiAuth, *ApiParams) *ApiResult) *ApiRouter {
    if handler == nil {
        return router.RemoveHandler(apiName)
    }
    router.handlersMap[apiName] = handler
    return nil
}

/*
RemoveHandler removes an api-handler.
*/
func (router *ApiRouter) RemoveHandler(apiName string) *ApiRouter {
    delete(router.handlersMap, apiName)
    return router
}

/*
GetAllHandlers returns all api-handlers as a map.
*/
func (router *ApiRouter) GetAllHandlers() map[string]func(*ApiContext, *ApiAuth, *ApiParams) *ApiResult {
    return router.handlersMap
}

/*
CallApi performs an API call.
*/
func (router *ApiRouter) CallApi(ctx *ApiContext, auth *ApiAuth, params *ApiParams) *ApiResult {
    apiName := ctx.GetApiName()
    handler := router.GetHandler(apiName)
    if handler == nil {
        return NewApiResult(StatusNotImplemented).SetMessage("No handler for API [" + apiName + "].")
    }
    return handler(ctx, auth, params)
}
