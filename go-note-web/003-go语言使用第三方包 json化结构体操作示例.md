go语言使用第三方包 json化结构体操作示例
 更新时间：2019年06月10日 11:59:22   作者：学习笔记666     我要评论

这篇文章主要介绍了go语言使用第三方包 json化结构体操作,结合实例形式分析了Go语言ffjson包git安装及结构体转json字符串相关操作技巧,需要的朋友可以参考下

本文实例讲述了go语言使用第三方包 json化结构体操作。分享给大家供大家参考，具体如下：

前提条件：安装好操作系统对应的git
```
go get -u github.com/pquerna/ffjson
```
-u参数：如果本地已经存在该包，则强制更新。

## 1、既然是把结构体转为json字符串，那么先来定义一个结构体

代码如下:
```
// 定义一个结构体
type NewsModel struct {
 Id int
 Title string
}
```

## 2、且看ffjson这包用什么方法来把结构体转为json字符串

代码如下:
```
func main()  {
 news := NewsModel{110,"hello"}
 res,err := ffjson.Marshal(news)
 if err != nil {
  fmt.Println("格式化错误")
  fmt.Println(err.Error())
  return
 }
 // 得到是字节数组，所以还有转为string
 fmt.Println(string(res))
}
```

打印：
```
{"Id":110,"Title":"hello"}
```
得到了一个json字符串

## 3、拓展

为结构体封装一个方法ToJson()专门来干这事
代码如下:
```
package main
import (
 "fmt"
 "github.com/pquerna/ffjson/ffjson"
)
// 定义一个结构体
type NewsModel struct {
 Id int
 Title string
}
// 定义一个方法
func (news NewsModel) ToJson() string  {
 res,err := ffjson.Marshal(news)
 if err != nil {
  return  err.Error()
 }
 // 得到是字节数组，所以还有转为string
 return string(res)
}
func main()  {
 news := NewsModel{110,"hello"}
 fmt.Println(news.ToJson()) // 打印：{"Id":110,"Title":"hello"}
}
```