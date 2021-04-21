package hello

import "fmt"

func hello(name string) string {
	return "Hello "+ name
}

func main() {
	fmt.Println(hello("Jason"))
}