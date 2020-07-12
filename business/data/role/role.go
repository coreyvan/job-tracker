package role

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// Add adds Role to database
func Add(ctx context.Context, gql *graphql.GraphQL, role Role) (Role, error) {
	c, err := add(ctx, gql, role)
	if err != nil {
		return Role{}, err
	}
	return c, nil
}

func add(ctx context.Context, gql *graphql.GraphQL, role Role) (Role, error) {
	if role.ID != "" {
		return Role{}, errors.New("Role contains id")
	}

	mutation, result := prepareAdd(role)
	if err := gql.Query(ctx, mutation, &result); err != nil {
		return Role{}, errors.Wrap(err, "failed to add Role")
	}

	if len(result.AddRole.Role) != 1 {
		return Role{}, errors.New("advisory id not returned")
	}

	role.ID = result.AddRole.Role[0].ID
	return role, nil
}

func prepareAdd(Role Role) (string, addResult) {
	var result addResult
	mutation := fmt.Sprintf(`
mutation {
	addRole(input: [{
		title: %q
		company: {
			id: %q
		}
		url: %q
		technologies: %q
		pay_lower: %d
		pay_upper: %d
		location: %q
		level: %q
		remote_possible: %t
		posted_on: %q
	}])
	%s
}`, Role.Title, Role.Company.ID, Role.URL,
		Role.Technologies, Role.PayLower, Role.PayUpper,
		Role.Location, Role.Level, Role.RemotePossible, Role.PostedOn.Format(time.RFC3339), result.document())

	return mutation, result
}

// GetOne retrieves a Role by ID
func GetOne(ctx context.Context, gql *graphql.GraphQL, id string) (Role, error) {
	c, err := getOne(ctx, gql, id)
	if err != nil {
		return Role{}, err
	}
	return c, nil
}

func getOne(ctx context.Context, gql *graphql.GraphQL, id string) (Role, error) {

	query := fmt.Sprintf(`
query {
	getRole(id: %q) {
		id
		title
		company {
			id
			name
		}
		url
		technologies
		pay_lower
		pay_upper
		location
		level
		remote_possible
		posted_on
	}
}`, id)

	var result struct {
		GetRole Role `json:"getRole"`
	}

	if err := gql.Query(ctx, query, &result); err != nil {
		return Role{}, errors.Wrap(err, "failed to find role")
	}

	if result.GetRole.ID == "" {
		return Role{}, errors.New("role not found")
	}

	return result.GetRole, nil
}

// GetOneByTitle retrieves a Role by searching by name
func GetOneByTitle(ctx context.Context, gql *graphql.GraphQL, query string) (Role, error) {
	c, err := getOneByTitle(ctx, gql, query)
	if err != nil {
		return Role{}, err
	}
	return c, nil
}

func getOneByTitle(ctx context.Context, gql *graphql.GraphQL, query string) (Role, error) {

	gquery := fmt.Sprintf(`
	query {
		queryRole(filter: { title: { anyofterms: %q } } ) {
			id
			title
			company {
				id
				name
			}
			url
			technologies
			pay_lower
			pay_upper
			location
			level
			remote_possible
			posted_on
		}
	}`, query)

	var result struct {
		GetRole []Role `json:"queryRole"`
	}

	if err := gql.Query(ctx, gquery, &result); err != nil {
		return Role{}, errors.Wrap(err, "failed to find role")
	}

	if len(result.GetRole) < 1 {
		return Role{}, errors.New("role not found")
	}

	if result.GetRole[0].ID == "" {
		return Role{}, errors.New("role not found")
	}

	return result.GetRole[0], nil
}

// List returns all companies
func List(ctx context.Context, gql *graphql.GraphQL, limit int) ([]Role, error) {
	gquery := fmt.Sprintf(`
 query {
	queryRole(first: %d) {
		id
		title
		company {
			id
			name
		}
		url
		technologies
		pay_lower
		pay_upper
		location
		level
		remote_possible
		posted_on
	}
}`, limit)

	var result struct {
		GetCompanies []Role `json:"queryRole"`
	}

	if err := gql.Query(ctx, gquery, &result); err != nil {
		return []Role{}, errors.Wrap(err, "failed to list companies")
	}

	return result.GetCompanies, nil
}

// Delete deletes a Role by ID
func Delete(ctx context.Context, gql *graphql.GraphQL, id string) error {
	if _, err := getOne(ctx, gql, id); err != nil {
		return errors.Wrap(err, "role does not exist")
	}

	if err := delete(ctx, gql, id); err != nil {
		return errors.Wrap(err, "deleting role")
	}

	return nil
}

func prepareDelete(RoleID string) (string, deleteResult) {
	var result deleteResult
	mutation := fmt.Sprintf(`
mutation {
	deleteRole(filter: { id: [%q] })
	%s
}`, RoleID, result.document())

	return mutation, result
}

func delete(ctx context.Context, gql *graphql.GraphQL, id string) error {
	if id == "" {
		return errors.New("missing role id")
	}

	mutation, result := prepareDelete(id)

	if err := gql.Query(ctx, mutation, &result); err != nil {
		return errors.Wrap(err, "failed to delete role")
	}

	if result.DeleteRole.NumUids == 0 {
		msg := fmt.Sprintf("failed to delete user: NumUids: %d  Msg: %s", result.DeleteRole.NumUids, result.DeleteRole.Msg)
		return errors.New(msg)
	}

	return nil
}
