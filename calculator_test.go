package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{b: 5, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 5, b: 4, want: 1},
		{a: 1, b: -1, want: 2},
		{a: -4, b: 5.5, want: -9.5},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 5, b: 4, want: 20},
		{a: 1, b: -1, want: -1},
		{a: -4},
		{a: -4, b: 4, want: -16},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{5, 2, 2.5, false},
		{25, 5, 5, false},
		{444, 0, 0, true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err == nil && tc.errExpected {
			t.Fatalf("Divide(%f, %f): error expected but got %f", tc.a, tc.b, got)
		} else if err != nil && !tc.errExpected {
			t.Fatalf("Divide(%f, %f): error not expected but got %v", tc.a, tc.b, err)
		} else if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f, error? %v", tc.a, tc.b, tc.want, got, err)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 16, want: 4, errExpected: false},
		{a: 64, want: 8, errExpected: false},
		{a: 0, want: 0, errExpected: false},
		{a: -16, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err == nil && tc.errExpected {
			t.Fatalf("Sqrt(%f): error expected but got %f", tc.a, got)
		} else if err != nil && !tc.errExpected {
			t.Fatalf("Sqrt(%f): error not expected but got %v", tc.a, err)
		} else if !tc.errExpected && tc.want != got {
			t.Errorf("Sqrt(%f): want %f, got %f, error? %v", tc.a, tc.want, got, err)
		}
	}
}
