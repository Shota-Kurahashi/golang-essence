package main

import (
	"basic-grammar/anys"
	"basic-grammar/array"
	"basic-grammar/basic"
	"basic-grammar/char"
	"basic-grammar/constructor"
	"basic-grammar/defers"
	forloop "basic-grammar/forLoop"
	"basic-grammar/interfaces"

	"basic-grammar/maps"
	"basic-grammar/pointer"
	"basic-grammar/structure"
)

func main() {
	basic.Main()
	forloop.Main()
	array.Main()
	char.Main()
	maps.Main()
	structure.Main()
	pointer.Main()
	anys.Main()
	constructor.NewUser("foo", 10)
	interfaces.Main()
	defers.Main()
}
