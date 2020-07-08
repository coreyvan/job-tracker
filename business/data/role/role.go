package role

import (
	"context"
	"fmt"

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
		Role.title: %q
		Role.company: %q
		Role.URL: %q
		Role.technologies: %q
		Role.pay_lower: %d
		Role.pay_upper: %d
		Role.location: %q
		Role.level: %q
		Role.remote_possible: %t
		Role.posted_on: %v
	}])
	%s
}`, Role.Title, Role.Company.ID, Role.URL,
		Role.Technologies, Role.PayLower, Role.PayUpper,
		Role.Location, Role.Level, Role.RemotePossible, Role.PostedOn, result.document())

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
	getRole(func: uid(%s)) {
		uid
		Role.name
		Role.description
		Role.industries
		Role.website
		Role.months
		Role.location
		Role.remote_possible
	}
}`, id)

	var result struct {
		GetRole []Role `json:"getRole"`
	}

	if err := gql.QueryPM(ctx, query, &result); err != nil {
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

// GetOneByName retrieves a Role by searching by name
func GetOneByName(ctx context.Context, gql *graphql.GraphQL, query string) (Role, error) {
	c, err := getOneByName(ctx, gql, query)
	if err != nil {
		return Role{}, err
	}
	return c, nil
}

func getOneByName(ctx context.Context, gql *graphql.GraphQL, query string) (Role, error) {

	gquery := fmt.Sprintf(`
 query {
	getRole(func:eq(Role.name, %q)) {
		uid
		Role.name
		Role.description
		Role.industries
		Role.website
		Role.months
		Role.location
		Role.remote_possible
	}
}`, query)

	var result struct {
		GetRole []Role `json:"getRole"`
	}

	if err := gql.QueryPM(ctx, gquery, &result); err != nil {
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
	getRoles(func: has(Role.name), first: %d) {
		uid
		Role.name
		Role.description
		Role.industries
		Role.website
		Role.months
		Role.location
		Role.remote_possible
	}
}`, limit)

	var result struct {
		GetCompanies []Role `json:"getRoles"`
	}

	if err := gql.QueryPM(ctx, gquery, &result); err != nil {
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
		return errors.Wrap(err, "failed to list companies")
	}
	fmt.Println(result.DeleteRole.NumUids)
	if result.DeleteRole.NumUids == 0 {
		msg := fmt.Sprintf("failed to delete user: NumUids: %d  Msg: %s", result.DeleteRole.NumUids, result.DeleteRole.Msg)
		return errors.New(msg)
	}

	return nil
}
