package http

import (
	"fmt"
	"strings"
)

type ConstraintViolation struct {
	Field   string
	Message string
}

type ConstraintViolationError struct {
	Violations []ConstraintViolation
}

func (e *ConstraintViolationError) AddViolation(field string, message string) {
	violation := ConstraintViolation{Field: field, Message: message}
	e.Violations = append(e.Violations, violation)
}

func (e *ConstraintViolationError) HasErrors() bool {
	return len(e.Violations) > 0
}

func (e *ConstraintViolationError) Error() string {
	var sb strings.Builder
	for _, violation := range e.Violations {
		sb.WriteString(fmt.Sprintf("Field: %s - Error: %s\n", violation.Field, violation.Message))
	}
	return sb.String()
}
