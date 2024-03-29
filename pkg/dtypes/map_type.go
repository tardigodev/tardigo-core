package dtypes

import (
	"fmt"
	"reflect"
)

type MapType struct{}

func (t MapType) GetName() string {
	return "map"
}

func (t MapType) Validate() error {
	return nil
}

func (t MapType) Convert(valueToConvert any) (any, error) {
	if reflect.TypeOf(valueToConvert).Kind() != reflect.Map {
		return nil, fmt.Errorf("failed to convert '%T' to map", valueToConvert)
	}
	return valueToConvert, nil
}
