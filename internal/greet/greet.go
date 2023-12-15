package greet

import "fmt"

func Hello(name string) string {
	return fmt.Sprint("Hello, ", name+"!")
}
