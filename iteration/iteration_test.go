package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	var tests = map[string]struct {
		character string
		frequency int
		expected  string
	}{
		"a,5":   {"a", 5, "aaaaa"},
		"b,5":   {"b", 5, "bbbbb"},
		"c,3":   {"c", 3, "ccc"},
		"bob,2": {"bob", 4, "bobbobbobbob"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := Repeat(test.character, test.frequency)
			if actual != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, actual)
			}

		})
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output:aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 5)
	}
}
