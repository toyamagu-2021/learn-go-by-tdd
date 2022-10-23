package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	characterIteration := 5
	for i := 0; i < b.N; i++ {
		Repeat("a", characterIteration)
	}
}

func ExampleRepeat() {
	iterated_chara := Repeat("a", 5)
	fmt.Println(iterated_chara)

	// Output: aaaaa
}
