本文实例讲述了Go语言集成mysql驱动、调用数据库、查询数据操作。分享给大家供大家参考，具体如下：

## 1、安装第三方mysql驱动包
```
go get -u github.com/go-sql-driver/mysql
```
## 2、连接数据库基本代码
代码如下:
```
package main
import (
        _"github.com/go-sql-driver/mysql"  // 注意前面的下划线_， 这种方式引入包只执行包的初始化函数
        "database/sql"
        "fmt"
)
func main()  {
        // 连接本地test数据库
        db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
        if err != nil {
                fmt.Println("连接数据库失败："+err.Error())
                return
        }
        // 查询user表
        rows,err := db.Query("select name,age from user")
        if err != nil {
                fmt.Println("查询错误："+err.Error())
                return
        }
        // 打印列名
        fmt.Println(rows.Columns()) // 打印：[name age] <nil>
}
```

## 4、循环扫描数据表取出数据

```
// for循环
for rows.Next()  {
        // 定义2个变量
        var name string
        var  age int
        // 扫描行并把扫描到到数据赋值
        rows.Scan(&name,&age)
        // 打印
        fmt.Println(name,age)
}
```
打印：
```
jack1 11
jack2 12
jack3 13
jack4 14
jack5 15
jack6 16
jack7 17
jack8 18
```