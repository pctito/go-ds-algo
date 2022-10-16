package main

import (
	"github.com/pctito/go-ds-algo/structs"
)

func main() {

	ll := new(structs.Linked)
	ll.PrintLs()
	ll.PrependInt(42)
	ll.PrependInt(5)
	ll.PrintLs()

}
