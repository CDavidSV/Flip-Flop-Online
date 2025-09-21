package validator

// FieldError represents a validation error for a specific field.
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"error"`
	Value   any    `json:"value,omitempty"`
}

// Validator holds validation errors.
type Validator struct {
	Errors []FieldError
}

// New creates a new Validator instance.
func New() *Validator {
	return &Validator{
		Errors: []FieldError{},
	}
}

// AddError adds an error to the errors slice.
func (v *Validator) AddError(field, message string, value any) {
	v.Errors = append(v.Errors, FieldError{
		Field:   field,
		Message: message,
		Value:   value,
	})
}

// Valid returns true if there are no validation errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// Check adds an error if the condition is false.
func (v *Validator) Check(ok bool, field, message string, value any) {
	if !ok {
		v.AddError(field, message, value)
	}
}
