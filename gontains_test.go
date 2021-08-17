package gontains

import (
	"testing"
)

func Test_Any(t *testing.T) {
	type TestData struct {
		name           string
		store          interface{}
		check          func(interface{}) bool
		expectedResult bool
	}
	data := []TestData{
		{
			"Simple Test",
			[]string{"test"},
			func(i interface{}) bool {
				return i.(string) == "test"
			},
			true,
		},
		{
			"Struct Test",
			[]TestData{
				{
					store: []string{"test2"},
				},
			},
			func(i interface{}) bool {
				return i.(TestData).store.([]string)[0] == "test2"
			},
			true,
		},
	}

	for _, dt := range data {
		result := Any(dt.store, dt.check)
		if result != dt.expectedResult {
			t.Errorf("Running Any() in %s was expected to return %v but returned %v", dt.name, dt.expectedResult, result)
		}
	}
}
