package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/company"
	"github.com/coreyvan/job-tracker/foundation/web"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type companyHandlers struct {
	gqlConfig data.GraphQLConfig
}

func (c *companyHandlers) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var compAdd company.Company
	if err := json.NewDecoder(r.Body).Decode(&compAdd); err != nil {
		return errors.Wrap(err, "decoding request body")
	}

	gql := data.NewGraphQL(c.gqlConfig)

	comp, err := company.Add(ctx, gql, compAdd)
	if err != nil {
		return errors.Wrap(err, "adding company")
	}
	return web.Respond(ctx, w, comp, http.StatusOK)
}

func (c *companyHandlers) getOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	gql := data.NewGraphQL(c.gqlConfig)

	comp, err := company.GetOne(ctx, gql, id)
	if err != nil {
		return errors.Wrap(err, "retrieving company")
	}

	return web.Respond(ctx, w, comp, http.StatusOK)
}

func (c *companyHandlers) search(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// vars := mux.Vars(r)
	query := r.URL.Query()["search"]

	gql := data.NewGraphQL(c.gqlConfig)

	comp, err := company.GetOneByName(ctx, gql, query[0])
	if err != nil {
		return errors.Wrap(err, "retrieving company")
	}

	return web.Respond(ctx, w, comp, http.StatusOK)
}
