package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/role"
	"github.com/coreyvan/job-tracker/foundation/web"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type roleHandlers struct {
	gqlConfig data.GraphQLConfig
}

func (ro *roleHandlers) getOne(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	gql := data.NewGraphQL(ro.gqlConfig)

	comp, err := role.GetOne(ctx, gql, id)
	if err != nil {
		return errors.Wrap(err, "retrieving role")
	}

	return web.Respond(ctx, w, comp, http.StatusOK)
}

func (ro *roleHandlers) search(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// vars := mux.Vars(r)
	query := r.URL.Query()["search"]

	gql := data.NewGraphQL(ro.gqlConfig)

	role, err := role.GetOneByTitle(ctx, gql, query[0])
	if err != nil {
		return errors.Wrap(err, "retrieving company")
	}

	return web.Respond(ctx, w, role, http.StatusOK)
}

func (ro *roleHandlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	gql := data.NewGraphQL(ro.gqlConfig)
	n := r.URL.Query()["limit"]
	limit, err := strconv.Atoi(n[0])
	if err != nil {
		return errors.Wrap(err, "invalid limit param")
	}

	role, err := role.List(ctx, gql, limit)
	if err != nil {
		return errors.Wrap(err, "retrieving company")
	}

	return web.Respond(ctx, w, role, http.StatusOK)
}

func (ro *roleHandlers) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var roleAdd role.Role
	if err := json.NewDecoder(r.Body).Decode(&roleAdd); err != nil {
		return errors.Wrap(err, "decoding request body")
	}

	gql := data.NewGraphQL(ro.gqlConfig)

	role, err := role.Add(ctx, gql, roleAdd)
	if err != nil {
		return errors.Wrap(err, "adding company")
	}
	return web.Respond(ctx, w, role, http.StatusOK)
}
