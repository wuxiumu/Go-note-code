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
loan float32
}
type Employee struct {
Human // 匿名字段
company string
money float32
}
//Human 实现 Sayhi 方法
func (h Human) SayHi() {
fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
//Human 实现 Sing 方法
func (h Human) Sing(lyrics string) {
fmt.Println("La la la la...", lyrics)

}
//Employee 重载 Human 的 SayHi 方法
func (e Employee) SayHi() {
fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
e.company, e.phone) //Yes you can split into 2 lines here.
}
// Interface Men 被 Human,Student 和 Employee 实现
//  因为这三个类型都实现了这两个方法
type Men interface {
SayHi()
Sing(lyrics string)
}
func main() {
mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}
// 定义 Men 类型的变量 i
var i Men
//i 能存储 Student
i = mike
fmt.Println("This is Mike, a Student:")
i.SayHi()
i.Sing("November rain")
//i 也能存储 Employee
i = Tom
fmt.Println("This is Tom, an Employee:")
i.SayHi()
i.Sing("Born to be wild")
// 定义了 slice Men
fmt.Println("Let's use a slice of Men and see what happens")
x := make([]Men, 3)
//T 这三个都是不同类型的元素，但是他们实现了 interface 同一个接口
x[0], x[1], x[2] = paul, sam, mike
for _, value := range x{
value.SayHi()
}
}