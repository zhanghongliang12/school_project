package main

import "fmt"

func main() {
	a := 10
	switch {
	case a > 10:
		fmt.Println(a)
	case a <= 10:
		fmt.Println("a<10,a is", a)

	}

}
