package integers

import "testing"

func TestAdder(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, sum, expected int) {
        t.Helper()
        if sum != expected {
            t.Errorf("expected '%d' but got '%d'", expected, sum)
        }
    }

    t.Run("2 + 2", func(t *testing.T) {
        sum := Add(2, 2)
        expected := 4

        if sum != expected {
            assertCorrectMessage(t, sum, expected)
        }
    })

    t.Run("10 + 2", func(t *testing.T) {
        sum := Add(10, 2)
        expected := 12

        if sum != expected {
            assertCorrectMessage(t, sum, expected)
        }
    })

    t.Run("-10 + 2", func(t *testing.T) {
        sum := Add(-10, 2)
        expected := -8

        if sum != expected {
            assertCorrectMessage(t, sum, expected)
        }
    })
}
