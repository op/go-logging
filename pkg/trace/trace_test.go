package trace

import (
	"testing"
)

func TestGetCallerFunctionName(t *testing.T){
	result := GetCallerFunctionName(0)
	expected := "TestGetCallerFunctionName"
	if result != expected {
		t.Errorf("actual=%q, expected=%q", result, expected)
	}
}

func TestGetCallerFunctionNameInsideLambda(t *testing.T) {
	func(){
		result := GetCallerFunctionName(0)
		expected := "TestGetCallerFunctionNameInsideLambda"
		if result != expected {
			t.Errorf("actual=%q, expected=%q", result, expected)
		}
	}()
}

func BenchmarkGetCallerFunctionName(t *testing.B){
	for i:=0; i < t.N; i++ {
		result := GetCallerFunctionName(0)
		expected := "BenchmarkGetCallerFunctionName"
		if result != expected {
			t.Errorf("actual=%q, expected=%q", result, expected)
		}
	}
}