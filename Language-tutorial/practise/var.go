package main

import "unsafe"
import "fmt"

const(
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)


func main(){
	println(a, b, c)
}