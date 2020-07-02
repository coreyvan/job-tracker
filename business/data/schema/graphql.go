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
	name: String! @search(by: [exact])
	description: String
	industries: [String]
	website: String
	months: Int
	location: String
	remote_possible: Boolean
}`
