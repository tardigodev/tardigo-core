package utils

import (
	"fmt"
	"reflect"

	"github.com/tardigodev/tardigo-core/pkg/dtypes"
)

func ConvertToRecord(anyRecord any, schema dtypes.Schema) (any, error) {
	reflectValue := reflect.ValueOf(anyRecord)
	if reflectValue.Kind() == reflect.Array || reflectValue.Kind() == reflect.Slice {
		if len(schema) != reflectValue.Len() {
			return nil, fmt.Errorf("record length %d not equal to schema length %d", reflectValue.Len(), len(schema))
		}

		record := make([]any, reflectValue.Len())
		for i := 0; i < reflectValue.Len(); i++ {
			convertedCol, err := schema[i].Type.Convert(reflectValue.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to convert column %s: with value %v -> %v", schema[i].Name, reflectValue.Index(i).Interface(), err)
			}
			record[i] = convertedCol
		}
		return record, nil
	}
	return nil, fmt.Errorf("record type %T not supported", anyRecord)
}
