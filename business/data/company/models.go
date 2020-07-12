package company

// Company represents the object model for a company
type Company struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description,omitempty"`
	Website        string   `json:"website,omitempty"`
	Industries     []string `json:"industries,omitempty"`
	Months         int      `json:"months,omitempty"`
	Location       string   `json:"location,omitempty"`
	RemotePossible bool     `json:"remote_possible,omitempty"`
}

// Ref reference to company
type Ref struct {
	UID string `json:"id"`
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

type deleteResult struct {
	DeleteCompany struct {
		Msg     string
		NumUids int
	} `json:"deleteCompany"`
}

func (deleteResult) document() string {
	return `{
		msg,
		numUids,
	}`
}
