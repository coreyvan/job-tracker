package data

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/ardanlabs/graphql"
	"github.com/coreyvan/job-tracker/business/data/ready"
	"github.com/coreyvan/job-tracker/business/data/schema"
	"github.com/pkg/errors"
)

// GraphQLConfig represents comfiguration needed to support managing, mutating,
// and querying the database.
type GraphQLConfig struct {
	URL            string
	AuthHeaderName string
	AuthToken      string
}

// Base base predicates that all dgraph nodes should have
type Base struct {
	ID        string    `json:"uid"`
	CreatedAt time.Time `json:"created_at"`
}

// NewGraphQL constructs a graphql value for use to access the databse.
func NewGraphQL(gqlConfig GraphQLConfig) *graphql.GraphQL {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	auth := graphql.WithAuth(gqlConfig.AuthHeaderName, gqlConfig.AuthToken)
	graphql := graphql.New(gqlConfig.URL, &client, auth)

	return graphql
}

// UpdateSchema creates/updates the schema for the database.
func UpdateSchema(gqlConfig GraphQLConfig, schemaConfig schema.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := ready.Validate(ctx, gqlConfig.URL, 5*time.Second)
	if err != nil {
		return errors.Wrapf(err, "waiting for database to be ready, database timed out or does not exist")
	}

	gql := NewGraphQL(gqlConfig)

	schema, err := schema.New(gql, schemaConfig)
	if err != nil {
		return errors.Wrapf(err, "preparing schema")
	}

	if err := schema.Create(ctx); err != nil {
		return errors.Wrapf(err, "creating schema")
	}

	return nil
}
