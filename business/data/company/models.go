package company

import "github.com/coreyvan/job-tracker/business/data"

// Company represents the object model for a company
type Company struct {
	data.Base
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Website        string   `json:"website"`
	Industries     []string `json:"industries"`
	Months         int      `json:"months"`
	Location       string   `json:"location"`
	RemotePossible bool     `json:"remote_possible"`
}

type addResult struct {
	AddCompany struct {
		Company []struct {
			ID string `json:"id"`
		} `json:"company"`
	} `json:"addCompany"`
}

func (addResult) document() string {
	return `{
			company {
					id
			}
		}`
}
