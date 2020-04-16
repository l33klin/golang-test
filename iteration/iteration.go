package iteration

import "fmt"

// Repeat a character 5 times
func Repeat(character string, repeat int) string {
	
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}

// ExampleRepeat
func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
