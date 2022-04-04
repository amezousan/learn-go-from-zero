package integer

import "testing"

func TestFourOps(t *testing.T) {
	assertNum := func(t testing.TB, got, expected int) {
		if got != expected {
			t.Errorf("got %d, want %d", got, expected)
		}
	}

	t.Run("Add: Plus + Plus become plus", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertNum(t, sum, expected)
	})

	t.Run("Add: Plus = Minus become zero", func(t *testing.T) {
		sum := Add(2, -2)
		expected := 0
		assertNum(t, sum, expected)
	})

	t.Run("Add: Plus < Minus become minus", func(t *testing.T) {
		sum := Add(2, -4)
		expected := -2
		assertNum(t, sum, expected)
	})

	t.Run("Sub: Plus - Minus become plus", func(t *testing.T) {
		sum := Subtract(2, -4)
		expected := 6
		assertNum(t, sum, expected)
	})

	t.Run("Multiply: Minus * Minus become plus", func(t *testing.T) {
		sum := Multiply(-2, -4)
		expected := 8
		assertNum(t, sum, expected)
	})

	t.Run("Divide: Minus / Minus become plus", func(t *testing.T) {
		sum := Divide(-8, -4)
		expected := 2
		assertNum(t, sum, expected)
	})
}
