package dfaregexp

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		s  string
		ok bool
	}{
		// invalid.
		{"", false},
		{" ", false},
		{"-", false},
		{"_", false},
		{"1", false}, // start with digit.
		{"0abc", false},
		{"9test", false},
		{"123", false},
		{"1_underscore", false},
		{"0-hyphen", false},

		// valid, since all start with letters.
		{"a", true},
		{"Z", true},
		{"abc", true},
		{"ABC", true},
		{"aBC123", true},
		{"MyClass", true},
		{"camelCase", true},

		// valid, can be prefixed with '_' or '-' and must follow by at least one char letter/digit.
		{"_a", true},
		{"_private", true},
		{"_123", true},
		{"_1", true},
		{"-a", true},
		{"-test", true},
		{"-123", true},
		{"-1", true},

		// valid, mix chars without consecutive symbol.
		{"my_var", true},
		{"my-var", true},
		{"my_var_123", true},
		{"a1b2c3", true},
		{"test-case-1", true},
		{"A_1-B_2", true},
		{"kebab-case-name", true},
		{"snake_case_name", true},
		{"MixedCase_with-all123", true},
		{"a_b_c", true},
		{"a-b-c", true},
		{"test_", true}, // underscore suffix is OK
		{"test_1", true},
		{"a_1-b_2", true},

		// invalid, consecutive symbols
		{"__init__", false},
		{"a__b", false},
		{"a--b", false},
		{"a_-b", false},
		{"a-_b", false},
		{"--flag", false},
		{"_-test", false},
		{"-_test", false},

		// invalid, has `-` suffix.
		{"test-", false},
		{"a-", false},
		{"my_var-", false},
		{"abc123-", false},

		// invalid, contains invalid characters.
		{"a b", false},
		{"test!", false},
		{"hello@world", false},
		{"path/to/file", false},
		{"a.b", false},
		{"var$name", false},
		{"test#1", false},
		{"a+b", false},
		{"name=value", false},
		{"hello\tworld", false},
		{"hello\nworld", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("s=%q ok=%v", tt.s, tt.ok), func(t *testing.T) {
			if Valid(tt.s) != tt.ok {
				t.Errorf("Valid(%q) != %v", tt.s, tt.ok)
			}
		})
	}
}

var benchCases = []struct {
	name string
	s    string
}{
	{"short_valid", "abc"},
	{"short_invalid", "1ab"},
	{"medium_valid", "my_variable_name_123"},
	{"medium_invalid", "my__invalid__name"},
	{"long_valid", "this_is-a_very-long_identifier-with_many-segments_123"},
	{"long_invalid", "this_is-a_very-long_identifier-with_many-segments-"},
	{"complex_valid", "MixedCase_with-kebab_and-snake_123"},
	{"complex_invalid", "MixedCase__with--consecutive"},
}

func BenchmarkValid(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			for range b.N {
				Valid(bc.s)
			}
		})
	}
}
