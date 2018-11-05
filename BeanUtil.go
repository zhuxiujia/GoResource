package Go_Resource

import (
	"reflect"
	"log"
)

var DefaultResourceContext = make(map[string]interface{})

func Register(beanName string,value interface{})  {
	var v = reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		panic("bean `" + beanName + "` must be a ptr!")
	}
	DefaultResourceContext[beanName]=value
}


//`bean:"***"`支持指针和实现interface的struct,指针（指针类型必须完全相同）struct(结构体可以完全一样，或者是继承实现某个接口的 struct)
//扫描 给所有加了 `bean:"***"` 注解的成员 设置beanMap name关联的对象
func AutoSetResourceProperty(beanMap map[string]interface{}, enableLog bool) {
	for _, bean := range beanMap {
		var v = reflect.ValueOf(bean)
		if v.Kind() != reflect.Ptr {
			panic("Bean: Bean must be a pointer")
		}
		v = v.Elem()
		ScanAndSetProperty(v, beanMap, false)
	}
	//第二次遍历
	for key, bean := range beanMap {
		var v = reflect.ValueOf(bean)
		if v.Kind() != reflect.Ptr {
			panic("Bean: Bean must be a pointer")
		}
		v = v.Elem()
		ScanAndSetProperty(v, beanMap, enableLog)
		beanMap[key] = bean
	}
}

func ScanAndSetProperty(v reflect.Value, beanMap map[string]interface{}, printInfo bool) (count int64) {
	var t = v.Type()
	if v.Kind() != reflect.Struct {
		return count
	}
	for i := 0; i < t.NumField(); i++ {
		var typeFieldItem = t.Field(i)
		var beanName = typeFieldItem.Tag.Get("resource")
		if beanName == "" {
			continue
		}
		var field = v.Field(i)
		var mapBeanInterface = beanMap[beanName]
		if mapBeanInterface == nil {
			if printInfo {
				log.Println("[resource] property ", v.String(),".", beanName, " not register!")
			}
			continue
		}
		if typeFieldItem.Type.Kind() == reflect.Struct {
			//loop scan bean feild
			var c = ScanAndSetProperty(field, beanMap, printInfo)
			count = count + c
		}
		var mapBeanValue = reflect.ValueOf(mapBeanInterface)
		if mapBeanValue.Kind() != reflect.Ptr {
			panic("[resource] property = " + v.String() +"."+ beanName + " must be a poiter")
		}
		mapBeanValue = mapBeanValue.Elem()
		field.Set(mapBeanValue)
		count++
	}
	return count
}
