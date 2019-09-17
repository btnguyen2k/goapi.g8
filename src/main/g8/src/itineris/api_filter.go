package itineris

import (
	"time"
)

/*
IApiFilter is pluggable component that is used to intercept API call and do some pre-processing, intercept result and do post-processing before returning to caller.
*/
type IApiFilter interface {
	Call(IApiHandler, *ApiContext, *ApiAuth, *ApiParams) *ApiResult
}

/*
BaseApiFilter is abstract implementation of IApiFilter
*/
type BaseApiFilter struct {
	apiRouter  *ApiRouter
	nextFilter IApiFilter
}

/*----------------------------------------------------------------------*/

/*
AddPerfInfoFilter adds the following data to the "debug" field of API's result:

    {
        "t"   : timestamp when the API was called (UNIX milliseconds),
        "tstr": timestamp as human-readable string,
        "d"   : API's execution duration (in nanoseconds),
        "c"   : server's total concurrent API calls
    }
*/
type AddPerfInfoFilter struct {
	*BaseApiFilter
}

/*
NewAddPerfInfoFilter creates a new AddPerfInfoFilter instance.
*/
func NewAddPerfInfoFilter(apiRouter *ApiRouter, nextFilter IApiFilter) *AddPerfInfoFilter {
	return &AddPerfInfoFilter{&BaseApiFilter{apiRouter: apiRouter, nextFilter: nextFilter}}
}

/*
Call implements IApiFilter.Call
*/
func (f *AddPerfInfoFilter) Call(handler IApiHandler, ctx *ApiContext, auth *ApiAuth, params *ApiParams) *ApiResult {
	now := time.Now()
	var apiResult *ApiResult
	if f.nextFilter != nil {
		apiResult = f.nextFilter.Call(handler, ctx, auth, params)
	} else {
		apiResult = handler(ctx, auth, params)
	}
	if apiResult != nil {
		debugData := map[string]interface{}{
			"t":    now.UnixNano() / 1000000, // convert to milliseconds
			"tstr": now.Format(time.RFC3339),
			"d":    time.Since(now).Nanoseconds(),
			"c":    f.apiRouter.GetConcurrency(),
		}
		apiResult.SetDebugInfo(debugData)
	}
	return apiResult
}

/*----------------------------------------------------------------------*/

/*
LoggingFilter performs logging before and after API call.
*/
type LoggingFilter struct {
	*BaseApiFilter
	logger IApiLogger
}

/*
NewLoggingFilter creates a new AddPerfInfoFilter instance.
*/
func NewLoggingFilter(apiRouter *ApiRouter, nextFilter IApiFilter, logger IApiLogger) *LoggingFilter {
	return &LoggingFilter{BaseApiFilter: &BaseApiFilter{apiRouter: apiRouter, nextFilter: nextFilter}, logger: logger}
}

/*
Call implements IApiFilter.Call
*/
func (f *LoggingFilter) Call(handler IApiHandler, ctx *ApiContext, auth *ApiAuth, params *ApiParams) *ApiResult {
	f.logger.PreApiCall(f.apiRouter.GetConcurrency(), ctx, auth, params)
	now := time.Now()
	var apiResult *ApiResult
	if f.nextFilter != nil {
		apiResult = f.nextFilter.Call(handler, ctx, auth, params)
	} else {
		apiResult = handler(ctx, auth, params)
	}
	f.logger.PostApiCall(time.Since(now).Nanoseconds(), f.apiRouter.GetConcurrency(), ctx, auth, params, apiResult)
	return apiResult
}

/*----------------------------------------------------------------------*/

/*
AuthenticationFilter performs authentication check before calling API.
*/
type AuthenticationFilter struct {
	*BaseApiFilter
	auth IApiAuthenticator
}

/*
NewAuthenticationFilter creates a new AuthenticationFilter instance.
*/
func NewAuthenticationFilter(apiRouter *ApiRouter, nextFilter IApiFilter, auth IApiAuthenticator) *AuthenticationFilter {
	return &AuthenticationFilter{BaseApiFilter: &BaseApiFilter{apiRouter: apiRouter, nextFilter: nextFilter}, auth: auth}
}

/*
Call implements IApiFilter.Call
*/
func (f *AuthenticationFilter) Call(handler IApiHandler, ctx *ApiContext, auth *ApiAuth, params *ApiParams) *ApiResult {
	if !f.auth.Authenticate(ctx, auth) {
		return ResultNoPermission
	}
	if f.nextFilter != nil {
		return f.nextFilter.Call(handler, ctx, auth, params)
	}
	return handler(ctx, auth, params)
}
