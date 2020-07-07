package gdal

import (
	"testing"
)

// TestDataTypes tries various functions for gdal types
func TestDataTypes(t *testing.T) {
	var b DataType
	b = Byte

	if ret := b.Size(); ret != 8 {
		t.Fail()
	}

	if b.IsComplex() != 0 {
		t.Fail()
	}
}
