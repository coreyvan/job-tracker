package commands

import (
	"log"

	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/schema"
	"github.com/pkg/errors"
)

// Schema updates the graphql schema
func Schema(gqlCfg data.GraphQLConfig, schemaCfg schema.Config, log *log.Logger) error {
	log.Println("schema: Starting schema update")
	defer log.Println("schema: Update finished")

	if err := data.UpdateSchema(gqlCfg, schemaCfg); err != nil {
		return errors.Wrap(err, "updating schema")
	}

	return nil
}
