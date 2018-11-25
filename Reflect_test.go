package Go_Resource

import (
	"reflect"
	"fmt"
	"testing"
)

type Reflect_A interface {
	Print() string
}
type Reflect_B struct {
}

func (this Reflect_B) Print() string {
	return "fuck"
}

func TestReflect(t *testing.T) {
	var a Reflect_A

	var b = Reflect_B{}

	var vb = reflect.ValueOf(b)
	var aType = reflect.TypeOf(&a)
	fmt.Println(aType.String())
	var newv = vb.Convert(aType.Elem())
	var newA = newv.Interface().(Reflect_A)
	fmt.Println(newA.Print())
}
