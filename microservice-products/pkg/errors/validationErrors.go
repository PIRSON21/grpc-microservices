package errors

type ValidationError interface {
	Field() string
	Error() string
}

type ValidationErrorImpl struct {
	field string
	err   error
}

// Error implements the error interface
func (v *ValidationErrorImpl) Error() string {
	return v.err.Error()
}

func (v *ValidationErrorImpl) Field() string {
	return v.field
}

// MarshalJSON customizes the JSON representation of ValidationError.
func (v *ValidationErrorImpl) MarshalJSON() ([]byte, error) {
	return []byte(`{"` + v.field + `": "` + v.err.Error() + `"}`), nil
}
