package iteration

import "testing"

func isEqual(t *testing.T, repeated, expected string) {
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("a x 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		isEqual(t, repeated, expected)
	})

	t.Run("0 x 5", func(t *testing.T) {
		repeated := Repeat("0", 5)
		expected := "00000"
		isEqual(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
