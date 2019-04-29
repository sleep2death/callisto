package callisto

import "strings"

// Hello is a test function
func Hello(name string) string {
	arr := []string{"Hello", name}
	return strings.Join(arr, ", ")
}
