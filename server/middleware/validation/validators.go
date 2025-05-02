package validation

import (
	"fmt"
)

func GreaterThan(value int, minValue int, fieldName string) error {
	if value <= minValue {
		return fmt.Errorf("invalid field %s, value %d, must be greater than %d", fieldName, value, minValue)
	}
	return nil
}
func GreaterOrEqual(value int, secondValue int, fieldName string) error {
	if value < secondValue {
		return fmt.Errorf("invalid field %s, value %d, must be greater or equal to %d", fieldName, value, secondValue)
	}
	return nil
}
