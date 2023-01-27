package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	actual, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	expected := int32(2)
	require.Equal(t, expected, actual)
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	actual, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	expected := int32(0)
	require.Equal(t, expected, actual)
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	actual, err := arith.Multiplication(2, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	expected := int32(4)
	require.Equal(t, expected, actual)
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	actual, err := arith.Division(4, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	expected := int32(2)
	require.Equal(t, expected, actual)
}
