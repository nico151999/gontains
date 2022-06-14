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


func TestGeneric(t *testing.T) {
	type args[T comparable] struct {
		store []T
		check T
	}
	tests := []struct {
		name string
		args args[string]
		want bool
	}{
		{
			name: "Test successful with string", args: args[string]{store: []string{"one", "two", "three"}, check: "two"}, want: true,
		},
		{
			name: "Test unsuccessful with string", args: args[string]{store: []string{"one", "two", "three"}, check: "four"}, want: false,
		},
		{
			name: "Test with empty store", args: args[string]{store: []string{}, check: "four"}, want: false,
		},
		{
			name: "Test with empty check", args: args[string]{store: []string{"one", "two", "three"}, check: ""}, want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generic(tt.args.store, tt.args.check); got != tt.want {
				t.Errorf("Generic() = %v, want %v", got, tt.want)
			}
		})
	}
}
