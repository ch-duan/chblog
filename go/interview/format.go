package interview

import "fmt"

type TestFormat struct {
	Name string
}

func (t *TestFormat) String() string {
	return "My name is " + t.Name
}

func (t *TestFormat) TestPrintf() {
	fmt.Println("My name is " + t.Name)
}

// 这里第二个打印结果会是{ch},因为传值定义的方法会默认生成指针方式的方法，而指针方式定义的方法不会生成传值的方式的方法（值方式的类型的String没有定义）
func FormatTestString() {
	t := TestFormat{}
	t.Name = "ch"
	fmt.Println(&t)
	fmt.Println(t)
	t.TestPrintf()

}
