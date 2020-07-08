package schema

// schema is a convenience variable for the schema and public key required for
// the project. This helps to support the training and quick start to getting
// the project up and running.
var schema = struct {
	document string
}{
	document: _document,
}

// _document represents the schema for the project. It is written to be processed
// by the Go templating engine to provide required parameters to be injected into
// the schema. In a production system the schema should be embedded using a tool
// like pakcr from the Buffalo project (https://github.com/gobuffalo/packr).
var _document = `
type Company {
	id: ID!
	name: String! @search(by: [term])
	description: String
	industries: [String] @search(by: [term])
	website: String
	months: Int
	location: String
	remote_possible: Boolean @search
}

type Role {
	id: ID!
	title: String! @search(by: [term])
	company: Company!
	url: String
	technologies: [String] @search(by: [term])
	pay_lower: Int
	pay_upper: Int
	location: String
	level: String
	remote_possible: Boolean @search
	posted_on: String
}
	`
