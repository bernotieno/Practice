## removeduplicate
### Instructions
Write a function that takes a slice of integers as an argument and returns the slice where all duplicate integers have been removed.

If the slice is empty, return the slice itself.

### Expected function
```
func RemoveDuplicate(slice []int) []int {

}
```

### Usage
Here is a possible program to test your function:
```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(RemoveDuplicate([]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(RemoveDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(RemoveDuplicate([]int{}))
}
```
And its output:
```
$ go run .
[1 2 3 4 5 6 7 8 9 10]
[1 2 3]
[]
```