package dtypes

import (
	"fmt"
	"strconv"
)

type IntegerType struct {
	IsUnsigned bool
	Byte       uint8
}

func (t IntegerType) GetName() string {
	return "integer"
}

func (t IntegerType) Validate() error {
	if t.Byte > 8 {
		return fmt.Errorf("byte size must be <= 8")
	}
	return nil
}

func (t IntegerType) Convert(valueToConvert any) (any, error) {
	var convFromString func(string) (any, error)
	if t.IsUnsigned {
		convFromString = func(v string) (any, error) {
			return strconv.ParseUint(v, 10, int(t.Byte)*8)
		}
	} else {
		convFromString = func(v string) (any, error) {
			return strconv.ParseInt(v, 10, int(t.Byte)*8)
		}
	}

	switch v := valueToConvert.(type) {
	case string:
		return convFromString(v)
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
		return nil, fmt.Errorf("failed to convert '%T' to integer", valueToConvert)
	}
}
