package main

import (
	"fmt"
)

type StructA struct {
	VarA int
	VarB string
	VarC error
}

type StructB struct {
	VarA StructA
}

func main() {

}
