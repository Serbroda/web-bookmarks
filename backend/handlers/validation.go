package handlers

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

func (e *ConstraintViolationError) Error() string {
	var sb strings.Builder
	for _, violation := range e.Violations {
		sb.WriteString(fmt.Sprintf("Field: %s - Error: %s\n", violation.Field, violation.Message))
	}
	return sb.String()
}

func validateUser(username string, age int) *ConstraintViolationError {
	var violations ConstraintViolationError

	if len(username) == 0 {
		violations.AddViolation("Username", "Username darf nicht leer sein")
	}
	if age < 18 {
		violations.AddViolation("Age", "Alter muss mindestens 18 sein")
	}

	if len(violations.Violations) > 0 {
		return &violations
	}
	return nil
}
