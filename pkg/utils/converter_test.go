package utils

import (
	"testing"

	"github.com/tardigodev/tardigo-core/pkg/dtypes"
)

func TestConvertToRecord(t *testing.T) {
	schema := dtypes.Schema{
		dtypes.Column{
			Name: "a",
			Type: dtypes.StringType{},
		},
		dtypes.Column{
			Name: "b",
			Type: dtypes.StringType{},
		},
		dtypes.Column{
			Name: "c",
			Type: dtypes.StringType{},
		},
	}
	convertedRecord, err := ConvertToRecord([]int{1, 2, 3}, schema)

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	record, ok := convertedRecord.([]any)

	if !ok {
		t.Errorf("Expected %T, got %T", []any{}, convertedRecord)
	}

	if len(record) != 3 {
		t.Errorf("Expected %d, got %d", 3, len(record))
	}

	if record[0] != "1" || record[1] != "2" || record[2] != "3" {
		t.Errorf("Expected %v, got %v", []any{"1", "2", "3"}, record)
	}

	_, err = ConvertToRecord(1, schema)

	if err != nil {
		expectedError := "record type int not supported"
		if err.Error() != expectedError {
			t.Errorf("Expected %s, got %s", expectedError, err.Error())
		}
	}

	_, err = ConvertToRecord([]any{1}, schema)

	if err != nil {
		expectedError := "record length 1 not equal to schema length 3"
		if err.Error() != expectedError {
			t.Errorf("Expected %s, got %s", expectedError, err.Error())
		}
	}

	_, err = ConvertToRecord([]any{[]any{1}, 2, 3}, schema)
	if err != nil {
		expectedError := "failed to convert column a: with value [1] -> failed to convert '[]interface {}' to string"
		if err.Error() != expectedError {
			t.Errorf("Expected %s, got %s", expectedError, err.Error())
		}
	}
}
