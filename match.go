package gohelpconvey

import (
	"fmt"
	"regexp"
)

const (
	success = ""
)

// ShouldMatch returns an empty string if the actual value matches all provided regexes
func ShouldMatch(actual interface{}, expected ...interface{}) string {
	if len(expected) < 1 {
		return "Missing regex to match"
	}

	actualString, ok := actual.(string)
	if ok == false {
		return fmt.Sprintf("'%v' is not an string", actual)
	}
	actualValue := []byte(actualString)

	for i := len(expected) - 1; i >= 0; i-- {
		pattern, ok := expected[i].(string)
		if ok == false {
			return fmt.Sprintf("'%v' is not an string", pattern)
		}
		ok, err := regexp.Match(pattern, actualValue)
		if err != nil {
			return fmt.Sprintf("'%v' failed: %v", pattern, err)
		}
		if ok == false {
			return fmt.Sprintf("Expected: '%s'\nActual:   '%s'", pattern, actualValue)
		}
	}

	return success
}

// ShouldNotMatch returns an empty string if the actual value does not match all provided regexes
func ShouldNotMatch(actual interface{}, expected ...interface{}) string {
	if len(expected) < 1 {
		return "Missing regex to match"
	}

	actualString, ok := actual.(string)
	if ok == false {
		return fmt.Sprintf("'%v' is not an string", actual)
	}
	actualValue := []byte(actualString)

	for i := len(expected) - 1; i >= 0; i-- {
		pattern, ok := expected[i].(string)
		if ok == false {
			return fmt.Sprintf("'%v' is not an string", pattern)
		}
		ok, err := regexp.Match(pattern, actualValue)
		if err != nil {
			return fmt.Sprintf("'%v' failed: %v", pattern, err)
		}
		if ok == true {
			return fmt.Sprintf("Expected: Anything but '%s'\nActual:   '%s'", pattern, actualValue)
		}
	}

	return success
}
