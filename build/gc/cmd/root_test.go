package cmd

import "testing"

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