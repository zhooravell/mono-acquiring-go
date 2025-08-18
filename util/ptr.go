// revive:disable:var-naming
package util

// Pointer returns a pointer to the input value.
// Example usage:
//
//	var x int = 42
//	ptr := Pointer(x) // ptr points to x
func Pointer[K any](val K) *K {
	return &val
}

// PointerValue returns a value of the input pointer.
// Example usage:
//
//	var x int = 42
//	val := Pointer(&x) // 42
func PointerValue[K any](val *K) (zero K) {
	if val == nil {
		return zero
	}

	return *val
}
