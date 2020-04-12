package integers

import "fmt"

// Add takes two integers and returns the sum of them
func Add(a, b int) int {
    sum := a + b
    return sum
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}