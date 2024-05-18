package functions

import "fmt"

func Max(a []int) {
	// Check if the input slice is empty
	if len(a) == 0 {
		fmt.Println(0) // If the slice is empty, return 0 as there is no maximum value
	}

	// Initialize max variable with the first element of the slice
	max := a[0]

	// Iterate over the elements of the slice
	for _, v := range a {
		// Compare each element with the current maximum
		if v > max {
			// If the current element is greater than the current maximum,
			// update the max variable to the value of the current element
			max = v
		}
	}

	// After iterating over all elements, return the maximum value found
	fmt.Println(max)
}
