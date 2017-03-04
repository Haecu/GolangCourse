package main

import "fmt"

func main() {
	var n1, n2, i, accum int64 = 1, 1, 0, 0
	for i < 4000000 {
		i = n1 + n2
		if i%2 == 0 {
			accum = accum + i
		}
		n1 = n2
		n2 = i
	}
	fmt.Println(accum)
}

/*
By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.
*/
