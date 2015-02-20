package main

import (
	"fmt"
	"testing"
)

func TestMethod1(t *testing.T) {
	PrintStart()
	Method1()
}

func TestMethod2(t *testing.T) {
	PrintStart()
	Method2()
}

func TestMethod3(t *testing.T) {
	PrintStart()
	Method3()
}

func TestMethod4(t *testing.T) {
	PrintStart()
	Method4()
}

func TestMethod5(t *testing.T) {
	PrintStart()
	Method5()
}

func TestMethod6(t *testing.T) {
	PrintStart()
	Method6()
}

func PrintStart() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
