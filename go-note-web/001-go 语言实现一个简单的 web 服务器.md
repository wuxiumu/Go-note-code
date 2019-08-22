# go 语言实现一个简单的 web 服务器

学习Go语言的一些感受,不一定准确。

假如发生战争，JAVA一般都是充当航母战斗群的角色。

一旦出动,就是护卫舰、巡洋舰、航母舰载机、预警机、电子战飞机、潜艇等等
浩浩荡荡,杀将过去。

(JVM,数十个JAR包,Tomcat中间件,SSH框架,各种配置文件...天生就是重量级的,专为大规模作战)

而GO语言更像F35战斗轰炸机

单枪匹马,悄无声息,投下炸弹然后走人。

专属轰炸机,空战也会一点点.

实在搞不定,就叫它大哥F22。

(GO是编译型语言,不需要依赖,不需要虚拟机,可以调用C代码并且它足够简单，却非常全面)

计划Go语言学习的知识点
1. 搭建Http服务
2. 连接数据库
3. 本地IO
4. 多线程
5. 网络
6. 调用本地命令
7. 调用C语言代码

easy.go
```
package main

import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.ListenAndServe(":8090", nil)
}
```

运行
```
go run easy.go
```

这样Go语言以不到50行代码，编译之后不到7M的可执行文件，就实现了一个简易的静态服务器。