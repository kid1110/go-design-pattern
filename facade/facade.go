package facade

import "fmt"

type SubSystemOne struct{}
type SubSystemTwo struct{}
type SubSystemThree struct{}

func (o *SubSystemOne) MethodOne() {
	fmt.Println("子系统方法一")
}

func (t *SubSystemTwo) MethodTwo() {
	fmt.Println("子系统方法二")
}
func (tr *SubSystemThree) MethodThree() {
	fmt.Println("子系统方法三")
}

//facade将三个不同的子方法系统给隐藏掉，外界只看到facade的方法
type Facade struct {
	one   SubSystemOne
	two   SubSystemTwo
	three SubSystemThree
}

func (f *Facade) MethodA() {
	f.one.MethodOne()
	f.two.MethodTwo()
	f.three.MethodThree()
}
