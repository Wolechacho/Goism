package hamming

import (
	"errors"
	"fmt"
	"strings"
)

//Distance calulate the difference between DNA strands and return the difference counter
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of strands not equal")
	}
	counter := 0
	arrayA := strings.Split(a, "")
	arrayB := strings.Split(b, "")
	fmt.Println(arrayA)
	fmt.Println(arrayB)
	for i := 0; i < len(a); i++ {
		if arrayA[i] != arrayB[i] {
			counter++
		}
	}
	return counter, nil
}
