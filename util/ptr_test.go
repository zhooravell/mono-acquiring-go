package util

import (
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	tests := map[string]any{
		"string":  "test string",
		"int":     123,
		"int64":   int64(123),
		"float32": float32(1.23),
		"float64": float64(1.23),
		"bool":    true,
	}
	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			got := Pointer(val)

			if reflect.ValueOf(got).Kind() != reflect.Ptr {
				t.Error("Pointer() should return a pointer")
			}

			if val != *got {
				t.Errorf("Pointer() = %v, want %v", got, val)
			}
		})
	}
}

func TestPointerValue(t *testing.T) {
	tests := map[string]any{
		"string":  "test string",
		"int":     123,
		"int64":   int64(123),
		"float32": float32(1.23),
		"float64": float64(1.23),
		"bool":    true,
	}
	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			ptr := Pointer(val)
			got := PointerValue(ptr)

			if reflect.ValueOf(got).Kind() == reflect.Ptr {
				t.Error("PointerValue() should return a value of pointer")
			}

			if got != val {
				t.Errorf("PointerValue() = %v, want %v", got, val)
			}
		})
	}
}

func TestPointerValue_nil(t *testing.T) {
	t.Run("nil pointer for int", func(t *testing.T) {
		var val *int
		result := PointerValue(val)
		var expected int // zero value

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nil pointer for string", func(t *testing.T) {
		var val *string
		result := PointerValue(val)
		var expected string // zero value

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nil pointer for bool", func(t *testing.T) {
		var val *bool
		result := PointerValue(val)
		var expected bool // zero value

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nil pointer for error", func(t *testing.T) {
		var val *error
		result := PointerValue(val)
		var expected error // zero value

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
