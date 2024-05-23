package main

import (
	"fmt"
	"time"

	functions "functions/utils"

	"github.com/01-edu/z01"
)

func main() {
	time.Sleep(1 * time.Second)
	functions.Alphamirror("a")

	time.Sleep(1 * time.Second)
	functions.Alternate()

	time.Sleep(1 * time.Second)
	functions.ChunkSlice([]int{123, 234, 635, 34, 45, 44, 33, 24, 8659, 56, 46}, 3)

	time.Sleep(1 * time.Second)
	functions.Countparams()

	time.Sleep(1 * time.Second)
	functions.Max([]int{23, 4, 27, 58, 47})

	time.Sleep(1 * time.Second)
	functions.Reduceint([]int{23, 4, 27, 58, 47})

	time.Sleep(1 * time.Second)
	functions.Revparam()

	time.Sleep(1 * time.Second)
	functions.Rot13("bn")

	time.Sleep(1 * time.Second)
	functions.Strlen("hello there, im batman")

	time.Sleep(1 * time.Second)
	functions.Firstparam()

	time.Sleep(1 * time.Second)
	fmt.Println(functions.PreviousPrime(7))

	time.Sleep(1 * time.Second)
	functions.Compare("flowers", "books")

	time.Sleep(1 * time.Second)
	functions.Doop()

	time.Sleep(1 * time.Second)
	functions.FoldInt(functions.Add, []int{1, 2, 3}, 96)

	time.Sleep(1 * time.Second)
	functions.Tabmult()

	time.Sleep(1 * time.Second)
	functions.Replace()

	time.Sleep(1 * time.Second)
	fmt.Println(functions.WordMatch("hello", "hdbdefiufblxydgsildo"))

	time.Sleep(1 * time.Second)
	z01.PrintRune((functions.Nrune("hello", 5)))
	z01.PrintRune('\n')

	time.Sleep(1 * time.Second)
	functions.FirstRune("love")

	time.Sleep(1 * time.Second)
	functions.LastRune("love")

	time.Sleep(1 * time.Second)
	functions.Lastword()
}
