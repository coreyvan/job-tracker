package role

import (
	"encoding/json"
	"time"

	"github.com/coreyvan/job-tracker/business/data/company"
)

var (
	acceptedTimeFormats = []string{"2006-01-02", time.RFC3339}
)

// Role represents the object model for a role
// TODO: remote Role from names
type Role struct {
	ID             string          `json:"id"`
	Title          string          `json:"title"`
	Company        company.Company `json:"company"`
	URL            string          `json:"url"`
	Technologies   []string        `json:"technologies"`
	PayLower       int             `json:"pay_lower"`
	PayUpper       int             `json:"pay_upper"`
	Location       string          `json:"location"`
	Level          string          `json:"level"`
	RemotePossible bool            `json:"remote_possible"`
	PostedOn       time.Time       `json:"posted_on"`
}

type addResult struct {
	AddRole struct {
		Role []struct {
			ID string `json:"id"`
		} `json:"role"`
	} `json:"addRole"`
}

func (addResult) document() string {
	return `{
			role {
					id
			}
		}`
}

type deleteResult struct {
	DeleteRole struct {
		Msg     string
		NumUids int
	} `json:"deleteRole"`
}

func (deleteResult) document() string {
	return `{
		msg,
		numUids,
	}`
}

// MarshalJSON custom marshaller for Roles
func (r *Role) MarshalJSON() ([]byte, error) {
	type Alias Role
	return json.Marshal(struct {
		PostedOn string `json:"posted_on"`
		*Alias
	}{
		PostedOn: r.PostedOn.Format(time.RFC3339),
		Alias:    (*Alias)(r),
	})
}
