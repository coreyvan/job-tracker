package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/application"
	"github.com/coreyvan/job-tracker/foundation/web"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type applicationHandlers struct {
	gqlConfig data.GraphQLConfig
}

func (a *applicationHandlers) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var applicationAdd application.Application
	if err := json.NewDecoder(r.Body).Decode(&applicationAdd); err != nil {
		return errors.Wrap(err, "decoding request body")
	}

	gql := data.NewGraphQL(a.gqlConfig)

	application, err := application.Add(ctx, gql, applicationAdd)
	if err != nil {
		return errors.Wrap(err, "adding application")
	}
	return web.Respond(ctx, w, application, http.StatusOK)
}

func (a *applicationHandlers) delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	gql := data.NewGraphQL(a.gqlConfig)

	err := application.Delete(ctx, gql, id)
	if err != nil {
		return errors.Wrap(err, "deleting role")
	}

	return web.Respond(ctx, w, nil, http.StatusOK)
}

func (a *applicationHandlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	gql := data.NewGraphQL(a.gqlConfig)
	n := r.URL.Query()["limit"]
	limit, err := strconv.Atoi(n[0])
	if err != nil {
		return errors.Wrap(err, "invalid limit param")
	}

	applications, err := application.List(ctx, gql, limit)
	if err != nil {
		return errors.Wrap(err, "retrieving applications")
	}

	return web.Respond(ctx, w, applications, http.StatusOK)
}

func (a *applicationHandlers) getOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	gql := data.NewGraphQL(a.gqlConfig)

	application, err := application.GetOne(ctx, gql, id)
	if err != nil {
		return errors.Wrap(err, "retrieving application")
	}

	return web.Respond(ctx, w, application, http.StatusOK)
}
