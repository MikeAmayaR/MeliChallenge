package service

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	testCases := []struct {
		name     string
		messages [][]string
		want     string
	}{
		// Test case 1
		{"Test 1", [][]string{
			{"", "este", "es", "un", "mensaje"},
			{"este", "", "un", "mensaje"},
			{"", "", "es", "", "mensaje"},
		}, "este es un mensaje"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetMessage(tc.messages...)
			if got != tc.want {
				t.Errorf("GetMessage() = %v, want %v", got, tc.want)
			}
		})
	}
}
