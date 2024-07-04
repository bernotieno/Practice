## split
### Instructions
Write a function that receives a string and a separator and returns a slice of strings that results of splitting the string s by the separator sep.

### Expected function
```
func Split(s, sep string) []string {

}
```
### Usage
Here is a possible program to test your function :
```
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
```
And its output :
```
$ go run .
[]string{"Hello", "how", "are", "you?"}
$
```