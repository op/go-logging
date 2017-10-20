package trace

import (
	"testing"
)

func TestGetCallerFunctionNameSkippingAnnonymous(t *testing.T){

	result := GetCallerFunctionNameSkippingAnnonymous(0)
	expected := "TestGetCallerFunctionNameSkippingAnnonymous"
	if result != expected {
		t.Errorf("actual=%q, expected=%q", result, expected)
	}
}

func BenchmarkGetCallerFunctionNameSkippingAnnonymous(t *testing.B){

	for i:=0; i < t.N; i++ {
		result := GetCallerFunctionNameSkippingAnnonymous(0)
		expected := "TestGetCallerFunctionNameSkippingAnnonymous"
		if result != expected {
			t.Errorf("actual=%q, expected=%q", result, expected)
		}
	}
}