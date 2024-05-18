package functions

import (
	"fmt"
	"os"
	"strconv"
)

func Doop() {
	// Check if the number of command-line arguments is exactly 4
	// (program name + 3 arguments). If not, exit the program.
	if len(os.Args) != 4 {
		return
	}

	// Parse the first argument as a 64-bit integer
	val1, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	// Store the second argument as the operator
	operator := os.Args[2]
	// Parse the third argument as a 64-bit integer
	val2, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	// If there are errors in parsing either the first or third argument,
	// exit the program.
	if err1 != nil || err2 != nil {
		return
	}

	// Initialize a variable to store the result of the operation
	var result int64

	// Perform the operation based on the operator
	switch operator {
	case "+":
		// Addition
		result = val1 + val2
	case "-":
		// Subtraction
		result = val1 - val2
	case "*":
		// Multiplication
		result = val1 * val2
	case "/":
		// Division - Check for division by zero
		if val2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = val1 / val2
	case "%":
		// Modulo - Check for modulo by zero
		if val2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = val1 % val2
	default:
		// If the operator is not recognized, exit the program
		return
	}

	// Print the result of the operation
	fmt.Println(result)
}
