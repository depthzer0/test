package main

import(
	"fmt"
	"unsafe"
)

func main(){

	var a int8
	var b int
	
	fmt.Printf("a Sizeof "); fmt.Println(unsafe.Sizeof(a))
	fmt.Printf("b Sizeof "); fmt.Println(unsafe.Sizeof(b))
	fmt.Printf("int32(1) Sizeof "); fmt.Println(unsafe.Sizeof(int32(1)))
	fmt.Printf("1 000 000 000 Sizeof "); fmt.Println(unsafe.Sizeof(1000000000))
}