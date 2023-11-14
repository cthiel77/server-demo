package error

// Parameters list of parameter
type Parameters []Parameter

// Parameter represents the rest-api error parameter
type Parameter struct {
	FieldName string `json:"fieldName"`
}
