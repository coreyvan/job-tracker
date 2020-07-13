package application

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// Add adds application to database
func Add(ctx context.Context, gql *graphql.GraphQL, application Application) (Application, error) {
	a, err := add(ctx, gql, application)
	if err != nil {
		return Application{}, err
	}
	return a, nil
}

func add(ctx context.Context, gql *graphql.GraphQL, application Application) (Application, error) {
	if application.ID != "" {
		return Application{}, errors.New("application contains id")
	}

	mutation, result := prepareAdd(application)
	if err := gql.Query(ctx, mutation, &result); err != nil {
		return Application{}, errors.Wrap(err, "failed to add application")
	}

	if len(result.AddApplication.Application) != 1 {
		return Application{}, errors.New("advisory id not returned")
	}

	application.ID = result.AddApplication.Application[0].ID
	return application, nil
}

func prepareAdd(application Application) (string, addResult) {
	var result addResult
	mutation := fmt.Sprintf(`
mutation {
	addApplication(input: [{
		role: {
			id: %q
		}
		applied_on: %q
	}])
	%s
}`, application.Role.ID, application.AppliedOn.Format(time.RFC3339), result.document())

	return mutation, result
}

// GetOne retrieves a company by ID
func GetOne(ctx context.Context, gql *graphql.GraphQL, id string) (Application, error) {
	c, err := getOne(ctx, gql, id)
	if err != nil {
		return Application{}, err
	}
	return c, nil
}

func getOne(ctx context.Context, gql *graphql.GraphQL, id string) (Application, error) {

	query := fmt.Sprintf(`
	query {
		getApplication(id: %q) {
			id
			role {
				id
				title
				company {
					id
					name
				}
				posted_on
			}
			applied_on
		}
	}`, id)

	var result struct {
		GetApplication Application `json:"getApplication"`
	}

	if err := gql.Query(ctx, query, &result); err != nil {
		return Application{}, errors.Wrap(err, "failed to find application")
	}

	if result.GetApplication.ID == "" {
		return Application{}, errors.New("application not found")
	}

	return result.GetApplication, nil
}

// List returns all companies
func List(ctx context.Context, gql *graphql.GraphQL, limit int) ([]Application, error) {
	gquery := fmt.Sprintf(`
	query {
		queryApplication(first: %d){
		  id
		  role {
			  id
			  title
			  company {
				  id
				  name
			  }
		  }
		  applied_on
		}
	  }`, limit)

	var result struct {
		GetApplications []Application `json:"queryApplication"`
	}

	if err := gql.Query(ctx, gquery, &result); err != nil {
		return []Application{}, errors.Wrap(err, "failed to list applications")
	}

	return result.GetApplications, nil
}

// Delete deletes a company by ID
func Delete(ctx context.Context, gql *graphql.GraphQL, id string) error {
	if _, err := getOne(ctx, gql, id); err != nil {
		return errors.Wrap(err, "application does not exist")
	}

	if err := delete(ctx, gql, id); err != nil {
		return errors.Wrap(err, "deleting application")
	}

	return nil
}

func prepareDelete(applicationID string) (string, deleteResult) {
	var result deleteResult
	mutation := fmt.Sprintf(`
mutation {
	deleteApplication(filter: { id: [%q] })
	%s
}`, applicationID, result.document())

	return mutation, result
}

func delete(ctx context.Context, gql *graphql.GraphQL, id string) error {
	if id == "" {
		return errors.New("missing application id")
	}

	mutation, result := prepareDelete(id)

	if err := gql.Query(ctx, mutation, &result); err != nil {
		return errors.Wrap(err, "failed to delete application")
	}
	fmt.Println(result.DeleteApplication.NumUids)
	if result.DeleteApplication.NumUids == 0 {
		msg := fmt.Sprintf("failed to delete application: NumUids: %d  Msg: %s", result.DeleteApplication.NumUids, result.DeleteApplication.Msg)
		return errors.New(msg)
	}

	return nil
}
