//go:build !dfa
// +build !dfa

package dfaregexp

import "regexp"

var pattern = regexp.MustCompile(`^([a-zA-Z]|[_-][a-zA-Z0-9])[a-zA-Z0-9]*([_-][a-zA-Z0-9]+)*_?$`)

func Valid(s string) bool {
	return pattern.MatchString(s)
}
