package main

import (

	"pctito/go-datastructs-algos/structs"
)

func main() {

	ll := new(structs.Linked)
	ll.PrintLs()
	ll.Prepend(5)
	ll.PrintLs()

}
