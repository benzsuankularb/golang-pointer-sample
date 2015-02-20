package main

import (
	"errors"
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

func Method1() {
	a := StructA{1, "Created", nil}
	b := ChangeStructA_Pointer(&a)
	fmt.Println("VarA,Before is Equal to After :", a.VarA == b.VarA)
	fmt.Println("VarB,Before is Equal to After :", a.VarB == b.VarB)
	fmt.Println("VarC,After is not nil :", a.VarC != nil)
}

func Method2() {
	a := StructA{1, "Created", nil}
	b := ChangeStructA(a)
	fmt.Println("VarA,Before is Equal to After :", a.VarA == b.VarA)
	fmt.Println("VarB,Before is Equal to After :", a.VarB == b.VarB)
	fmt.Println("VarC,After is not nil :", a.VarC != nil)
}

func Method3() {
	a := StructA{1, "Created", nil}
	b := &StructA{1, "Created", nil}
	a = ChangeStructA_Pointer(&a)
	b = ChangeStructA_Pointer_ReturnPointer(b)
}

func ChangeStructA_Pointer(r *StructA) StructA {
	r.VarA = 2
	r.VarB = "Changed"
	r.VarC = errors.New("Error Changed.")
	return *r
}

func ChangeStructA_Pointer_ReturnPointer(r *StructA) *StructA {
	r.VarA = 2
	r.VarB = "Changed"
	r.VarC = errors.New("Error Changed.")
	return r
}

func ChangeStructA(r StructA) StructA {
	r.VarA = 2
	r.VarB = "Changed"
	r.VarC = errors.New("Error Changed.")
	return r
}
