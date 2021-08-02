package gontains

import "reflect"

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

// ContainsInt returns true if slice of integers s contains an integer k
func ContainsInt(s []int, k int) bool {
	for _, e := range s {
		if e == k {
			return true
		}
	}
	return false
}
