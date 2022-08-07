package main

import (
	"fmt"
	"github.com/col3name/algo/cmd/performacnce/pointer-vs-stack/stat"
	"unsafe"
)

type Employee struct {
	ID     int
	Name   string
	Age    int16
	Gender string
	Active bool
}

type EmployeeOptimized struct {
	Name   string
	Gender string
	ID     int
	Age    int16
	Active bool
}

func main() {
	var e Employee
	fmt.Println(fmt.Sprintf("Size of %T struct: %d bytes", e, unsafe.Sizeof(e)))
	var eOp EmployeeOptimized
	fmt.Println(fmt.Sprintf("Size of %T struct: %d bytes", eOp, unsafe.Sizeof(eOp)))
	var s string
	fmt.Println(fmt.Sprintf("Size of %T struct: %d bytes", s, unsafe.Sizeof(s)))
	s = "HelloHelloHellod3asdfasdfasdfasd"
	fmt.Println(fmt.Sprintf("Size of %T struct: %d bytes", s, unsafe.Sizeof(s)), len(s))
	stat.PrintMemUsage()
}