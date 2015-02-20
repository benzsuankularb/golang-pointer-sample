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

func PrintStart() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
