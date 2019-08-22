package main
import (
        "github.com/go-sql-driver/mysql"  // 注意前面的下划线_， 这种方式引入包只执行包的初始化函数
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