package commands

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/company"
	"github.com/pkg/errors"
)

// Import imports test data
func Import(gqlCfg data.GraphQLConfig, log *log.Logger) error {
	log.Println("import: Starting import")
	defer log.Println("import: Finished ")

	f, err := os.Open("test_data/companies.json")
	if err != nil {
		return errors.Wrap(err, "opening json file")
	}

	log.Printf("import: Reading file %s", f.Name())

	var companies []company.Company

	if err = json.NewDecoder(f).Decode(&companies); err != nil {
		return errors.Wrap(err, "decoding data")
	}

	gql := data.NewGraphQL(gqlCfg)

	for _, v := range companies {
		c, err := company.Add(context.Background(), gql, v)
		if err != nil {
			return errors.Wrap(err, "creating company")
		}

		log.Printf("import: created company %s id: %s", c.Name, c.ID)
	}
	return nil
}
