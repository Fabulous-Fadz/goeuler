package multiplesof3and5

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input int
	want  int
}{
	{input: 5, want: 3},
	{input: 10, want: 23},
	{input: 1_000, want: 233_168},
}

func TestRunBrute(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Brute-%d", tc.input), func(t *testing.T) {
			have := RunBrute(tc.input)
			assert.Equal(t, tc.want, have, fmt.Sprintf("Incorrect result. Want: %v, have: %v", tc.want, have))
		})
	}
}

func TestDivisibleBy(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		target   int
		expected int
	}{
		{"totalDivisibleBy-3-under-10", 3, 10, 18},
		{"totalDivisibleBy-7-under-1000", 7, 1_000, 71_071},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := totalDivisibleBy(tc.input, tc.target)
			assert.Equal(t, tc.expected, actual, "Check your formula")
		})
	}
}

func TestRunSmart(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Smart-%d", tc.input), func(t *testing.T) {
			have := RunSmart(tc.input)
			assert.Equal(t, tc.want, have, fmt.Sprintf("Incorrect results. Want: %v, have: %v", tc.want, have))
		})
	}

}

func BenchmarkBrute(b *testing.B) {
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("BenchBrute-%d", tc.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RunBrute(tc.input)
			}
		})
	}
}

func BenchmarkSmart(b *testing.B) {
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("BenchSmart-%d", tc.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RunSmart(tc.input)
			}
		})
	}
}
