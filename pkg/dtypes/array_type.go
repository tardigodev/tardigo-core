package dtypes

import (
	"fmt"
	"reflect"
)

type ArrayType struct{}

func (t ArrayType) GetName() string {
	return "array"
}

func (t ArrayType) Validate() error {
	return nil
}

func (t ArrayType) Convert(valueToConvert any) (any, error) {
	if reflect.TypeOf(valueToConvert).Kind() != reflect.Slice {
		return nil, fmt.Errorf("failed to convert '%T' to array", valueToConvert)
	}
	return valueToConvert, nil
}
