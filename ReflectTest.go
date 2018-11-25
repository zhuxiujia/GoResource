package main

import (
	"reflect"
	"fmt"
)

type A interface {
	Print() string
}
type B struct {
}

func (this B) Print() string {
	return "fuck"
}

func main() {
	var a A

	var b = B{}

	var vb = reflect.ValueOf(b)
	var aType = reflect.TypeOf(&a)
	fmt.Println(aType.String())
	var newv = vb.Convert(aType.Elem())
	var newA = newv.Interface().(A)
	fmt.Println(newA.Print())
}
