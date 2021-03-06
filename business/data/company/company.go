package company

import (
	"context"
	"fmt"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// Add adds company to database
func Add(ctx context.Context, gql *graphql.GraphQL, company Company) (Company, error) {
	c, err := add(ctx, gql, company)
	if err != nil {
		return Company{}, err
	}
	return c, nil
}

func add(ctx context.Context, gql *graphql.GraphQL, company Company) (Company, error) {
	if company.ID != "" {
		return Company{}, errors.New("company contains id")
	}

	mutation, result := prepareAdd(company)
	if err := gql.Query(ctx, mutation, &result); err != nil {
		return Company{}, errors.Wrap(err, "failed to add company")
	}

	if len(result.AddCompany.Company) != 1 {
		return Company{}, errors.New("advisory id not returned")
	}

	company.ID = result.AddCompany.Company[0].ID
	return company, nil
}

func prepareAdd(company Company) (string, addResult) {
	var result addResult
	mutation := fmt.Sprintf(`
mutation {
	addCompany(input: [{
		name: %q
		description: %q
		website: %q
		industries: %q
		months: %d
		location: %q
		remote_possible: %t
	}])
	%s
}`, company.Name, company.Description, company.Website,
		company.Industries, company.Months, company.Location,
		company.RemotePossible, result.document())

	return mutation, result
}

// GetOne retrieves a company by ID
func GetOne(ctx context.Context, gql *graphql.GraphQL, id string) (Company, error) {
	c, err := getOne(ctx, gql, id)
	if err != nil {
		return Company{}, err
	}
	return c, nil
}

func getOne(ctx context.Context, gql *graphql.GraphQL, id string) (Company, error) {

	query := fmt.Sprintf(`
	query {
		getCompany(id: %q) {
			id
			name
			description
			industries
			website
			months
			location
			remote_possible
		}
	}`, id)

	var result struct {
		GetCompany Company `json:"getCompany"`
	}

	if err := gql.Query(ctx, query, &result); err != nil {
		return Company{}, errors.Wrap(err, "failed to find company")
	}

	if result.GetCompany.ID == "" {
		return Company{}, errors.New("company not found")
	}

	return result.GetCompany, nil
}

// GetOneByName retrieves a company by searching by name
func GetOneByName(ctx context.Context, gql *graphql.GraphQL, query string) (Company, error) {
	c, err := getOneByName(ctx, gql, query)
	if err != nil {
		return Company{}, err
	}
	return c, nil
}

func getOneByName(ctx context.Context, gql *graphql.GraphQL, query string) (Company, error) {

	gquery := fmt.Sprintf(`
 query {
	queryCompany(filter: { name: { anyofterms: %q } }) {
		id
    	name
   		description
    	industries
    	website
    	months
    	location
    	remote_possible
	}
}`, query)

	var result struct {
		GetCompany []Company `json:"queryCompany"`
	}

	if err := gql.Query(ctx, gquery, &result); err != nil {
		return Company{}, errors.Wrap(err, "failed to find company")
	}

	if len(result.GetCompany) < 1 {
		return Company{}, errors.New("company not found")
	}

	if result.GetCompany[0].ID == "" {
		return Company{}, errors.New("company not found")
	}

	return result.GetCompany[0], nil
}

// List returns all companies
func List(ctx context.Context, gql *graphql.GraphQL, limit int) ([]Company, error) {
	gquery := fmt.Sprintf(`
	query {
		queryCompany(first: %d){
		  id
		  name
		  description
		  industries
		  website
		  months
		  location
		  remote_possible
		}
	  }`, limit)

	var result struct {
		GetCompanies []Company `json:"queryCompany"`
	}

	if err := gql.Query(ctx, gquery, &result); err != nil {
		return []Company{}, errors.Wrap(err, "failed to list companies")
	}

	return result.GetCompanies, nil
}

// Delete deletes a company by ID
func Delete(ctx context.Context, gql *graphql.GraphQL, id string) error {
	if _, err := getOne(ctx, gql, id); err != nil {
		return errors.Wrap(err, "company does not exist")
	}

	if err := delete(ctx, gql, id); err != nil {
		return errors.Wrap(err, "deleting company")
	}

	return nil
}

func prepareDelete(companyID string) (string, deleteResult) {
	var result deleteResult
	mutation := fmt.Sprintf(`
mutation {
	deleteCompany(filter: { id: [%q] })
	%s
}`, companyID, result.document())

	return mutation, result
}

func delete(ctx context.Context, gql *graphql.GraphQL, id string) error {
	if id == "" {
		return errors.New("missing company id")
	}

	mutation, result := prepareDelete(id)

	if err := gql.Query(ctx, mutation, &result); err != nil {
		return errors.Wrap(err, "failed to list companies")
	}
	fmt.Println(result.DeleteCompany.NumUids)
	if result.DeleteCompany.NumUids == 0 {
		msg := fmt.Sprintf("failed to delete user: NumUids: %d  Msg: %s", result.DeleteCompany.NumUids, result.DeleteCompany.Msg)
		return errors.New(msg)
	}

	return nil
}
