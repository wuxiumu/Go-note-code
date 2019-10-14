package main
import "fmt"
type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human // 匿名字段
	school string
}
type Employee struct {
	Human // 匿名字段
	company string
}
// 在 human 上面定义了一个 method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}

// method继承
// 前面一章我们学习了字段的继承，那么你也会发现Go的一个神奇之处，method也是可以继承的。如果匿名字段实现了
// 一个method，那么包含这个匿名字段的struct也能调用该method。让我们来看下面这个例子