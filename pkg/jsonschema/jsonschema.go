package jsonschema

type DataType string

const (
	DataTypeObject  DataType = "object"
	DataTypeNumber  DataType = "number"
	DataTypeInteger DataType = "integer"
	DataTypeString  DataType = "string"
	DataTypeArray   DataType = "array"
	DataTypeNull    DataType = "null"
	DataTypeBoolean DataType = "boolean"
)

type Definition struct {
	Type        DataType              `json:"type,omi