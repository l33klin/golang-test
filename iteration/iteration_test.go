package iteration

import "testing"

func TestRepeat(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, repeated, expected string) {
        t.Helper()
        if repeated != expected {
            t.Errorf("expected '%s' but got '%s'", expected, repeated)
        }
    }

    t.Run("Repeat four", func(t *testing.T) {
        repeated := Repeat("a", 4)
        expected := "aaaa"
        assertCorrectMessage(t, repeated, expected)
    })

    t.Run("Repeat one", func(t *testing.T) {
        repeated := Repeat("a", 1)
        expected := "a"
        assertCorrectMessage(t, repeated, expected)
    })

    t.Run("Repeat 10", func(t *testing.T) {
        repeated := Repeat("a", 10)
        expected := "aaaaaaaaaa"
        assertCorrectMessage(t, repeated, expected)
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}