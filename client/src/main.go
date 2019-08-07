package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/ssoroka/demo/contracts"
)

func main() {
	p := contracts.Person{
		Name: "Phil",
		Age:  35,
	}

	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	js.Global().Call("updateDOM", "Hello, World! "+string(b))
}
