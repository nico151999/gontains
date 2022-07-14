package gontains

import (
	"reflect"
	"strings"
)

// Contains returns true if a store (slice or array) contains a specified value.
func Contains(store interface{}, val interface{}) bool {
	switch reflect.TypeOf(store).Kind() {
	case reflect.Slice,
		reflect.Array:
		s := reflect.ValueOf(store)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == val {
				return true
			}
		}
	}
	return false
}

// ContainsString returns true if slice of strings s contains a string k
func ContainsString(s []string, k string) bool {
	for _, e := range s {
		if e == k {
			return true
		}
	}
	return false
}

// ContainsStringCaseInsensitive returns true if slice of strings s contains a string k while checking case insensitive
func ContainsStringCaseInsensitive(s []string, k string) bool {
	for _, e := range s {
		if strings.ToLower(e) == strings.ToLower(k) {
			return true
		}
	}
	return false
}

// ContainsInt returns true if slice of integers s contains an integer k
func ContainsInt(s []int, k int) bool {
	for _, e := range s {
		if e == k {
			return true
		}
	}
	return false
}

// Any returns true if a store (slice or array) contains an element that fulfills the condition defined in the check function
func Any(store interface{}, check func(interface{}) bool) bool {
	switch reflect.TypeOf(store).Kind() {
	case reflect.Slice,
		reflect.Array:
		s := reflect.ValueOf(store)
		for i := 0; i < s.Len(); i++ {
			if check(s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func Generic[T comparable](store []T, check T) bool {
	for _, element := range store {
		if element == check {
			return true
		}
	}
	return false
}

// Can be a custom function or a builtin function like strings.EqualFold
type CompareFn[T comparable] func(a T, b T) bool

// Allows passing a custom function to compare with. GenericFn will return when compareFn returns true.
func GenericFn[T comparable](store []T, check T, compareFn CompareFn[T]) bool {
	for _, element := range store {
		if compareFn(element, check) {
			return true
		}
	}
	return false
}
