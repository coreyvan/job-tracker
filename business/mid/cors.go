package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/coreyvan/job-tracker/foundation/web"
)

// CORS manages CORS response
func CORS(log *log.Logger) web.Middleware {
	// This is the actual middleware function to be executed.
	m := func(before web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			_, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			// Set headers for CORS
			// TODO: this should be configurable from the config of the app
			w.Header().Set("Access-Control-Allow-Headers:",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")

			if r.Method == "OPTIONS" {
				return web.Respond(ctx, w, nil, http.StatusOK)
			}

			// Return the error so it can be handled further up the chain.
			return before(ctx, w, r)
		}

		return h
	}

	return m
}
