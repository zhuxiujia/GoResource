package Go_Resource

import (
	"testing"
	"fmt"
)

type PrintName interface {
	PrintName() string
}

type B struct {
	Name string
}

func (this B) PrintName() string {
	return this.Name
}

type A struct {
	B        PrintName `resource:"b"`         //使用`resource:"b"` 注解注入对象 Annotated injection object
	BPointer *B        `resource:"b_pointer"` //使用`resource:"b"` 注解注入指针对象
}

//验证注解是否有效
func (a A) Print() {
	fmt.Println(a.B.PrintName())
}

//验证注解指针是否有效
func (a A) PrintBPointer() {
	fmt.Println(a.BPointer.Name)
}

func Test_Inject(t *testing.T) {
	var b = B{
		Name: "is B!",
	}
	var a A
	var bAddress = &b

	//You can also customize using your own context, as long as the context type meets the map[string]interface{}
	//注册到context中,值必须为指向对象的指针类型
	//你也可以自定义使用自己的context,只要context类型满足map[string]interface{}
	Register("b", &b)
	Register("a", &a)
	Register("b_pointer", &bAddress)

	//Automatically set all contain ` resource: "*" ` properties
	//自动设置所有包含`resource:"*"`的属性
	AutoSetResourceProperty(DefaultResourceContext, true)

	//validation
	//验证
	a.Print()
	a.PrintBPointer()
}
