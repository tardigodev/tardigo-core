package dtypes

import (
	"fmt"
)

type StringType struct{}

func (t StringType) GetName() string {
	return "string"
}

func (t StringType) Validate() error {
	return nil
}

func (t StringType) Convert(valueToConvert any) (any, error) {
	convToString := func(any) (string, error) {
		return fmt.Sprintf("%v", valueToConvert), nil
	}
	switch v := valueToConvert.(type) {
	case string:
		return v, nil
	case float32, float64:
		return convToString(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return convToString(v)
	case bool:
		return convToString(v)
	case []byte:
		return string(v), nil
	case []rune:
		return string(v), nil
	default:
		return nil, fmt.Errorf("failed to convert '%T' to string", valueToConvert)
	}
}
