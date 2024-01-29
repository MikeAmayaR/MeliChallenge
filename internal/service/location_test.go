package service

import (
	"math"
	"testing"
)

func TestGetLocation(t *testing.T) {
	testCases := []struct {
		name         string
		distances    []float32
		wantX        float32
		wantY        float32
		expectingErr bool
	}{
		// Test case 1: Utilizar distancias conocidas que resulten en una posición exacta
		{"Test 1", []float32{100.0, 115.5, 142.7}, 42.0, 32.0, false},
		{"Test Number of Distances", []float32{100.0, 115.5}, 0.0, 0.0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotX, gotY, err := GetLocation(tc.distances...)

			if tc.expectingErr {
				if err == nil {
					t.Errorf("Expected an error for test case '%v' but didn't get one", tc.name)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error for test case '%v' but got one: %v", tc.name, err)
				} else if math.Abs(float64(gotX-tc.wantX)) > 0.001 || math.Abs(float64(gotY-tc.wantY)) > 0.001 {
					t.Errorf("GetLocation() = (%v, %v), want (%v, %v)", gotX, gotY, tc.wantX, tc.wantY)
				}
			}
		})
	}
}
