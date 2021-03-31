// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/superorbital/wordchain/models"
	"github.com/superorbital/wordchain/restapi/operations"
	"github.com/superorbital/wordchain/restapi/operations/wordchain"
	types "github.com/superorbital/wordchain/types"
	words "github.com/superorbital/wordchain/words"
)

//go:generate swagger generate server --target ../../wordchain --name WordChains --spec ../swagger.yaml --principal interface{} --exclude-main

// updateItem updates a single todo
func getChain(body *models.ModelsWordchainPrefs) (string, error) {

	// FIXME: We should use the defaults passed in at the command line.
	prefs := types.Preferences{
		WordFile: "",
		Length:   5,
		Divider:  "-",
		Prepend:  "",
		Postpend: "",
		Seed:     "",
		Type:     []string{"adjective", "noun"},
	}

	if body.Length != 0 {
		prefs.Length = body.Length
	}
	if body.Divider != "" {
		prefs.Divider = body.Divider
	}
	if body.Postpend != "" {
		prefs.Postpend = body.Postpend
	}
	if body.Prepend != "" {
		prefs.Prepend = body.Prepend
	}
	if body.Seed != "" {
		prefs.Seed = body.Seed
	}
	if body.Type != nil {
		prefs.Type = body.Type
	}

	chain, err := words.Random(prefs)
	if err != nil {
		return "", err
	}

	result := types.Chain{Chain: chain}
	json, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(json), nil

}

func configureFlags(api *operations.WordChainsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WordChainsAPI) http.Handler {
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

	//if api.WordchainGetRandomHandler == nil {
	api.WordchainGetRandomHandler = wordchain.GetRandomHandlerFunc(func(params wordchain.GetRandomParams) middleware.Responder {
		if chain, err := getChain(params.Body); err != nil {
			return wordchain.NewGetRandomDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		} else {
			return wordchain.NewGetRandomOK().WithPayload(chain)
		}
	})
	//}

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
	return handler
}
