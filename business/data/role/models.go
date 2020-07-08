package role

import (
	"encoding/json"
	"time"

	"github.com/coreyvan/job-tracker/business/data/company"
	"github.com/pkg/errors"
)

var (
	acceptedTimeFormats = []string{"2006-01-02", time.RFC3339}
)

// Role represents the object model for a role
type Role struct {
	ID             string          `json:"Role.ID"`
	Title          string          `json:"Role.title"`
	Company        company.Company `json:"Role.company"`
	URL            string          `json:"Role.URL"`
	Technologies   []string        `json:"Role.technologies"`
	PayLower       int             `json:"Role.pay_lower"`
	PayUpper       int             `json:"Role.pay_upper"`
	Location       string          `json:"Role.location"`
	Level          string          `json:"Role.level"`
	RemotePossible bool            `json:"Role.remote_possible"`
	PostedOn       time.Time       `json:"Role.posted_at"`
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

// UnmarshalJSON custom unmarshaler for Role
func (r *Role) UnmarshalJSON(data []byte) error {
	v := make(map[string]interface{})

	if err := json.Unmarshal(data, &v); err != nil {
		return errors.Wrap(err, "unmarshalling Role")
	}

	var ok bool
	for _, t := range acceptedTimeFormats {
		postedTime, err := time.Parse(t, v["posted_at"].(string))
		if err == nil {
			ok = true
			r.PostedOn = postedTime
			break
		}
	}

	if !ok {
		r.PostedOn = time.Now()
	}

	if id, ok := v["uid"]; ok {
		r.ID = id.(string)
	}

	if _, ok := v["title"]; !ok {
		return errors.New("role title does not exist")
	}
	r.Title = v["title"].(string)

	r.Company = v["company"].(company.Company)
	return nil
}
