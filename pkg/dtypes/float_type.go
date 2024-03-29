package dtypes

import (
	"fmt"
	"strconv"
)

type FloatType struct {
	Byte uint8
}

func (t FloatType) GetName() string {
	return "float"
}

func (t FloatType) Validate() error {
	if t.Byte > 8 {
		return fmt.Errorf("byte size must be <= 8")
	}
	return nil
}

func (t FloatType) Convert(valueToConvert any) (any, error) {
	convFromString := func(v string) (any, error) {
		return strconv.ParseFloat(v, int(t.Byte)*8)
	}

	switch v := valueToConvert.(type) {
	case string:
		return convFromString(v)
	case float32, float64:
		return convFromString(fmt.Sprintf("%v", v))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return convFromString(fmt.Sprintf("%v", v))
	case bool:
		if v {
			return convFromString("1")
		} else {
			return convFromString("0")
		}
	case []byte:
		return convFromString(string(v))
	case []rune:
		return convFromString(string(v))
	default:
		return nil, fmt.Errorf("failed to convert '%T' to float", valueToConvert)
	}
}
