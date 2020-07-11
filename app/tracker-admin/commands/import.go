package commands

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/ardanlabs/graphql"
	"github.com/coreyvan/job-tracker/business/data"
	"github.com/coreyvan/job-tracker/business/data/company"
	"github.com/coreyvan/job-tracker/business/data/role"
	"github.com/pkg/errors"
)

// Import imports test data
func Import(gqlCfg data.GraphQLConfig, log *log.Logger, subCommand string) error {
	log.Println("import: Starting import")
	defer log.Println("import: Finished ")

	gql := data.NewGraphQL(gqlCfg)
	switch subCommand {
	case "company":
		_, err := importCompanies("test_data/companies.json", gql)
		if err != nil {
			return errors.Wrap(err, "importing companies")
		}
	case "role":
		_, err := importRoles("test_data/roles.json", gql)
		if err != nil {
			return errors.Wrap(err, "importing roles")
		}
	}
	return nil
}

func importCompanies(file string, gql *graphql.GraphQL) ([]company.Company, error) {
	var companies []company.Company

	f, err := os.Open(file)
	if err != nil {
		return companies, errors.Wrap(err, "opening json file")
	}

	if err = json.NewDecoder(f).Decode(&companies); err != nil {
		return []company.Company{}, errors.Wrap(err, "decoding data")
	}

	var retComp []company.Company
	for _, v := range companies {
		c, err := company.Add(context.Background(), gql, v)
		if err != nil {
			return []company.Company{}, errors.Wrap(err, "creating company")
		}
		retComp = append(retComp, c)
	}
	return retComp, nil
}

func importRoles(file string, gql *graphql.GraphQL) ([]role.Role, error) {
	var roles []role.Role

	// b, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	return roles, errors.Wrap(err, "opening json file")
	// }

	f, err := os.Open(file)
	if err != nil {
		return roles, errors.Wrap(err, "opening json file")
	}

	if err = json.NewDecoder(f).Decode(&roles); err != nil {
		return []role.Role{}, errors.Wrap(err, "decoding data")
	}
	log.Print(roles)

	var retRole []role.Role
	for _, v := range roles {
		log.Println(v)
		r, err := role.Add(context.Background(), gql, v)
		if err != nil {
			return []role.Role{}, errors.Wrap(err, "creating role")
		}
		retRole = append(retRole, r)
	}
	return retRole, nil
}
