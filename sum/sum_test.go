package sum

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectMessage(t, got, want, numbers)
	})

}
