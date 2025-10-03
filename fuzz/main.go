package main

import "fmt"

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	fmt.Println(Reverse("hello world"))
}
