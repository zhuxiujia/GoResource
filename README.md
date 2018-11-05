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
 type EventServiceImpl struct {
	  SmsEventHandler SmsEventHandler `resource:"SmsEventHandler"` //这里添加resource注解
 }
</pre>
第四步，提供需要注解的对象和被注解的对象
<pre>
 Go_Resource.Register("EventServiceImpl",&EventServiceImpl{})
 Go_Resource.Register("SmsEventHandler",&SmsEventHandler{})
 utils.AutoSetResourceProperty(Go_Resource.DefaultResourceContext,true)//自动设置属性为注解对象
</pre>
测试代码请查看https://github.com/zhuxiujia/Go-Resource/blob/master/BeanUtil_test.go
