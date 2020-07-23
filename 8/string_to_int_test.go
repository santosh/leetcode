package stringtoint

import "testing"

func TestTrim(t *testing.T) {
	t.Run("trim spaces at left", func(t *testing.T) {
		got := trimLeft("    ")
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("eliminate all non digit after last digit", func(t *testing.T) {
		got := stripRight("45 rast 654")
		want := "45"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
func TestStringToInt(t *testing.T) {
	t.Run("test positive integer", func(t *testing.T) {
		got := myAtoi("    654    ")
		want := 654

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test negitive integer", func(t *testing.T) {
		got := myAtoi("   -100")
		want := -100

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test negitive overflow", func(t *testing.T) {
		got := myAtoi("  -4294967297")
		want := -2147483648

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test positive overflow", func(t *testing.T) {
		got := myAtoi("4294967297")
		want := 2147483647

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
