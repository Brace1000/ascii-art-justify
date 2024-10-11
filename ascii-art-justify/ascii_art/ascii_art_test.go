package ascii_art

import (
	"testing"
)

// TestColorPicker tests the ColorPicker function to ensure it returns the correct ANSI color codes
func TestColorPicker(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected string
	}{
		{name: "Reset Color", color: "reset", expected: "\u001b[39m"},
		{name: "Red Color", color: "red", expected: "\u001b[31m"},
		{name: "Green Color", color: "green", expected: "\u001b[32m"},
		{name: "Magenta Color", color: "magenta", expected: "\u001b[35m"},
	}

	// Loop through each test case
	// Call the ColorPicker function with the test color
	// Check if the result matches the expected output
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := ColorPicker(tc.color)

			if result != tc.expected {
				t.Errorf("ColorPicker(%s) returned %s, expected %s", tc.color, result, tc.expected)
			}
		})
	}
}

// TestGetCi tests the getCi function to ensure it finds the correct index and length of the substring
// Check if the starting index matches the expected value
// Check if the length of the substring matches the expected value

// func TestGetCi(t *testing.T) {
// 	text := "hello world"
// 	substring := "world"
// 	ci := colorMap(substring, text)

// 	if ci != 6 {
// 		t.Errorf("getCi returned index %d, expected 6", ci)
// 	}

// 	if n != len(substring) {
// 		t.Errorf("getCi returned length %d, expected %d", n, len(substring))
// 	}
// }
