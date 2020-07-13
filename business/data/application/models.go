package application

import (
	"time"

	"github.com/coreyvan/job-tracker/business/data/role"
)

// Application represents an application to a job role
type Application struct {
	ID        string    `json:"id"`
	Role      role.Role `json:"role"`
	AppliedOn time.Time `json:"applied_on"`
}

type addResult struct {
	AddApplication struct {
		Application []struct {
			ID string `json:"id"`
		} `json:"application"`
	} `json:"addApplication"`
}

func (addResult) document() string {
	return `{
			application {
				id
			}
		}
	}`
}

type deleteResult struct {
	DeleteApplication struct {
		Msg     string
		NumUids int
	} `json:"deleteApplication"`
}

func (deleteResult) document() string {
	return `{
		msg,
		numUids,
	}`
}
