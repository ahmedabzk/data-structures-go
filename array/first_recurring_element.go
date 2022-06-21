// Golang program to find the first recurring element in an array


package main

import "fmt"

func main() {
	var arr [6]int
	var item int = 0
	var index int = 0

	fmt.Printf("Enter array elements: \n")
	for i := 0; i <= 5; i++ {
		// fmt.Printf("Elements: arr[%d]: ", i)
		fmt.Scanf("%d", &arr[i])
	}

	index = -1
	for i := 0; i < 5; i++ {
		for j := i + 1; j <= 5; j++ {
			if arr[i] == arr[j] {
				item = arr[j]
				index = j
				goto OUT
			}
		}
	}

OUT:
	if index != -1 {
		fmt.Printf("%d repeated @ %d index", item, index)
	} else {
		fmt.Printf("There is no repeated element")
	}
}
