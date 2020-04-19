package array

import "testing"

func TestRepeat(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want int, numbers [5]int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    }

    t.Run("Repeat four", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

        assertCorrectMessage(t, got, want, numbers)
    })
}

// func BenchmarkRepeat(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         Repeat("a", 5)
//     }
// }