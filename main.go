package main

import (
	"errors"
	"fmt"
)

type StructA struct {
	VarA int
	VarB string
	VarC error // A pointer of error
}

type StructB struct {
	VarA StructA
	VarB *StructA
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
	str1 := fmt.Sprintf("%p", &a)
	str2 := fmt.Sprintf("%p", b)
	fmt.Println("Address of A and B is Equal :", str1 == str2)
}

func Method4() {
	a := &StructA{1, "Created", nil}
	addressOfPointer := fmt.Sprintf("%p", &a)
	pointer := fmt.Sprintf("%p", a)
	objectA := fmt.Sprint("%p", *a)
	fmt.Println("Address Of Pointer Of A :", addressOfPointer)
	fmt.Println("Pointer Point To Address :", pointer)
	fmt.Println("Object Of Address That Pointer Point To", objectA)
}

func Method5() {
	a := &StructA{1, "Created", errors.New("Created Error")}
	b := Copy(a)
	fmt.Println("Address A And B is difference :", fmt.Sprintf("%p", a), fmt.Sprintf("%p", &b))
	fmt.Println("VarB Address of A and B is difference :", fmt.Sprintf("%p", &a.VarB), fmt.Sprintf("%p", &b.VarB))
	fmt.Println("VarC Address of A and B is difference :", fmt.Sprintf("%p", &a.VarC), fmt.Sprintf("%p", &b.VarC))
	fmt.Println("VarC of A and B point to same address :", fmt.Sprintf("%p", a.VarC), fmt.Sprintf("%p", b.VarC))
}

func Method6() {
	a := StructA{1, "Created", errors.New("Created Error")}
	b1 := StructB{ a, &a }
	b2 := CopyB(&b1)
	fmt.Println("Address StructB B1 And B2 is difference :", fmt.Sprintf("%p", &b1), fmt.Sprintf("%p", &b2))
	fmt.Println("Address VarB in VarA of B1 And B2 is difference :", fmt.Sprintf("%p", &b1.VarA.VarB), fmt.Sprintf("%p", &b2.VarA.VarB))
	fmt.Println("Address VarB of B1 And B2 is difference :", fmt.Sprintf("%p", &b1.VarB), fmt.Sprintf("%p", &b2.VarB))
	fmt.Println("Address that VarB of B1 and B2 pointed is same :", fmt.Sprintf("%p", b1.VarB), fmt.Sprintf("%p", b2.VarB))
}

func CopyB(r *StructB) StructB {
	return *r
}

func Copy(r *StructA) StructA {
	// Notice
	// r is pointer point to an address.
	// *r is object that address holding.
	return *r
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
