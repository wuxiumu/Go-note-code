package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)
func sayHelloName(w http.ResponseWriter, r *http.Request) {
    // 解析url传递的参数
    r.ParseForm()
    //在服务端打印信息
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("Scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        // join() 方法用于把数组中的所有元素放入一个字符串。
        // 元素是通过指定的分隔符进行分隔的
        fmt.Println("val:", strings.Join(v, ""))
    }
    // 输出到客户端
    fmt.Fprintf(w, "hello astaxie!")
}
func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        // 执行解析模板
        // func (t *Template) Execute(wr io.Writer, data interface{}) error {
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}
func main() {
    //设置访问路由
    http.HandleFunc("/", sayHelloName)
    http.HandleFunc("/login", login)
    //设置监听端口
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndserve:", err)
    }
}