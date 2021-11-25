package cmd

import (
	"testing"
)

func TestVersionsAreEqual(t *testing.T) {
	var tests = []struct {
		input    []string
		expected bool
	}{
		{[]string{"17.0.0", "17.0.0"}, true},
		{[]string{"17.2.1", "17.2.1"}, true},
		{[]string{"1.1.0", "1.2.0"}, false},
		{[]string{"11.1.0", "12.1.0"}, false},
	}

	for _, test := range tests {
		got := versionsAreEqual(test.input[0], test.input[1])
		if got != test.expected {
			t.Errorf("TEST FAILED - (%v, %v) - Expected: %v, Got: %v", test.input[0], test.input[1], test.expected, got)
		}
	}
}

func TestGetProfileName(t *testing.T) {
	var tests = []struct {
		input    []string
		expected string
	}{
		{[]string{}, "DEFAULT"},
		{[]string{"mock3", "mock4"}, "DEFAULT"},
		{[]string{"--profile"}, "DEFAULT"},
		{[]string{"-p"}, "DEFAULT"},
		{[]string{"-p="}, "DEFAULT"},
		{[]string{"--profile="}, "DEFAULT"},
		{[]string{"--profile=", "mockybalboa"}, "DEFAULT"},
		{[]string{"double", "--profile", "mock_name"}, "mock_name"},
		{[]string{"-p", "mock2", "double"}, "mock2"},
		{[]string{"double", "-p=mock5", "double"}, "mock5"},
		{[]string{"--profile=mock6"}, "mock6"},
	}

	for _, test := range tests {
		got := getProfileName(test.input)
		if got != test.expected {
			t.Errorf("TEST FAILED - Expected: %v, Got: %v", test.expected, got)
		}
	}
}