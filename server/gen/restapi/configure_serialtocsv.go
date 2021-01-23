// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"local.packages/gen/restapi/serialtocsv"
	"local.packages/gen/restapi/serialtocsv/common"
	"local.packages/handler"
	"local.packages/mymiddleware"
)

//go:generate swagger generate server --target ../../gen --name Serialtocsv --spec ../../../swagger/serial-csv-converter.v1.yaml --api-package serialtocsv --principal interface{}

func configureFlags(api *serialtocsv.SerialtocsvAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *serialtocsv.SerialtocsvAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CommonPutStartHandler = &handler.PutStartHandler{}
	api.CommonPutStopHandler = &handler.PutStopHandler{}
	api.CommonGetStatusHandler = &handler.GetStatusHandler{}
	api.CommonGetDataHandler = &handler.GetDataHandler{}
	api.CommonPutResetHandler = &handler.PutResetHandler{}
	api.CommonPutSendHandler = &handler.PutSendHandler{}

	if api.CommonGetDataHandler == nil {
		api.CommonGetDataHandler = common.GetDataHandlerFunc(func(params common.GetDataParams) middleware.Responder {
			return middleware.NotImplemented("operation common.GetData has not yet been implemented")
		})
	}
	if api.CommonGetStatusHandler == nil {
		api.CommonGetStatusHandler = common.GetStatusHandlerFunc(func(params common.GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation common.GetStatus has not yet been implemented")
		})
	}
	if api.CommonPutResetHandler == nil {
		api.CommonPutResetHandler = common.PutResetHandlerFunc(func(params common.PutResetParams) middleware.Responder {
			return middleware.NotImplemented("operation common.PutReset has not yet been implemented")
		})
	}
	if api.CommonPutStartHandler == nil {
		api.CommonPutStartHandler = common.PutStartHandlerFunc(func(params common.PutStartParams) middleware.Responder {
			return middleware.NotImplemented("operation common.PutStart has not yet been implemented")
		})
	}
	if api.CommonPutStopHandler == nil {
		api.CommonPutStopHandler = common.PutStopHandlerFunc(func(params common.PutStopParams) middleware.Responder {
			return middleware.NotImplemented("operation common.PutStop has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return mymiddleware.Recovery(mymiddleware.AccessLog(handler))
}
