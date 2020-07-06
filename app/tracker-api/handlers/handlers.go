package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/mid"
	"github.com/coreyvan/job-tracker/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, gql data.GraphQLConfig, log *log.Logger) http.Handler {

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log))

	t := trackerHandlers{}
	app.Handle("GET", "/", nil, t.home)

	c := companyHandlers{gqlConfig: gql}
	app.Handle("POST", "/company", nil, c.create)
	app.Handle("GET", "/company", []string{"search", "{query}"}, c.search)
	app.Handle("GET", "/company/{id}", nil, c.getOne)

	return app.Mux()
}

type trackerHandlers struct{}

func (t *trackerHandlers) home(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.Respond(ctx, w, nil, http.StatusOK)
}
