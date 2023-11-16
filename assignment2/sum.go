package main

import (
	"fmt"
)

func main() {

	sum := 0
	for count := 1; count <= 100; count = count + 1 {
		if count%2 == 0 {
			sum += count
		} else {
			sum -= count
		}
	}
	fmt.Println(sum)

}
