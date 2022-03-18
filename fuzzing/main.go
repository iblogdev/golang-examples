package main

import "fmt"

func main() {
	fmt.Println(Reverse("Hello"))
}

func Reverse(s string) string {
	//out := []byte(s)
	out := []rune(s)
	for i, j := 0, len(out)-1; i < len(out)/2; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}
