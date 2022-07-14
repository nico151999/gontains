package gontains

import (
	"fmt"
	"strings"
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

func TestGenericFn(t *testing.T) {
	type args struct {
		store     []string
		check     string
		compareFn CompareFn[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Successful test with not equal casing ",
			args: args{store: []string{"Some.One@mail.com", "Someone.Else@mail.com", "Third.Person@mail.com"}, check: "third.person@mail.com", compareFn: strings.EqualFold},
			want: true,
		},
		{
			name: "Successful test with equal casing ",
			args: args{store: []string{"Some.One@mail.com", "Someone.Else@mail.com", "third.person@mail.com"}, check: "third.person@mail.com", compareFn: strings.EqualFold},
			want: true,
		},
		{
			name: "Unsuccessful test with lowercase custom comparer ",
			args: args{store: []string{"Some.One@mail.com", "Someone.Else@mail.com"}, check: "third.person@mail.com", compareFn: strings.EqualFold},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenericFn(tt.args.store, tt.args.check, tt.args.compareFn); got != tt.want {
				t.Errorf("GenericFn() = %v, want %v", got, tt.want)
			}
		})
	}

}

func ExampleGenericFn() {
	// Custom string comparer for ignoring casing
	fmt.Println(GenericFn([]string{"Some.One@mail.com", "Someone.Else@mail.com"}, "third.person@mail.com", strings.EqualFold))
	//Output: false

}
