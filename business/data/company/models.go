package company

import "github.com/coreyvan/job-tracker/business/data"

// Company represents the object model for a company
type Company struct {
	data.Base
	Name           string   `json:"Company.name"`
	Description    string   `json:"Company.description"`
	Website        string   `json:"Company.website"`
	Industries     []string `json:"Company.industries"`
	Months         int      `json:"Company.months"`
	Location       string   `json:"Company.location"`
	RemotePossible bool     `json:"Company.remote_possible"`
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
