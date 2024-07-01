package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1:]

	for _, words := range input {
		if words == "01" || words == "galaxy" || words == "galaxy 01" {
			for _, c := range "Alert!!!" {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
	}
}

// Stella's solution
/*func main() {
	Comcheck := map[string]bool{
		"01":        true,
		"galaxy":    true,
		"galaxy 01": true,
	}
	for _, v := range os.Args[1:] {
		if Comcheck[v] {
			fmt.Println("Alert!!!")
		}
	}
}
*/
