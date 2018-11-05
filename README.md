# Go-Resource
# a smart easy to use Golang inject  framework,like spring mvc
# 一个简单易用的golang注解框架

#使用方法
第一步
<pre>
go get github.com/zhuxiujia/Go-Resource
</pre>
第二步
<pre>
 import ("github.com/zhuxiujia/Go-Resource")
</pre>
第三步，定义属性，添加注解
<pre>
type B struct {
	Name string
}
 type A struct {
	B        B  `resource:"b"`         //使用`resource:"b"` 注解注入对象 Annotated injection object
	BPointer *B `resource:"b_pointer"` //使用`resource:"b"` 注解注入指针对象
}
</pre>
第四步，提供需要注解的对象和被注解的对象
<pre>
 Go_Resource.Register("b", &B{})
 Go_Resource.Register("a", &A{})
 
 bAddress := &B{}
 Go_Resource.Register("b_pointer", &bAddress)
 utils.AutoSetResourceProperty(Go_Resource.DefaultResourceContext,true)//自动设置属性为注解对象
</pre>
测试代码请查看https://github.com/zhuxiujia/Go-Resource/blob/master/BeanUtil_test.go
