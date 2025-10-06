package errors

type ValidationError struct {
	Field string
	Err   error
}

// Error implements the error interface
func (v ValidationError) Error() string {
	return v.Field + ": " + v.Err.Error()
}

// MarshalJSON customizes the JSON representation of ValidationError.
func (v ValidationError) MarshalJSON() ([]byte, error) {
	return []byte(`{"` + v.Field + `": "` + v.Err.Error() + `"}`), nil
}
