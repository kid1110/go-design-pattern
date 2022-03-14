package builder

import "fmt"

//builder抽象类
type PersonBuilder interface {
	BuildHead()
	BuildBody()
	BuildArm()
	BuildLeg()
}

//具体抽象类实现
type PersonThinBuilder struct {
	PersonBuilder
}

func (p *PersonThinBuilder) BuildHead() {
	fmt.Println("瘦子的头")
}
func (p *PersonThinBuilder) BuildArm() {
	fmt.Println("瘦子的手臂")
}

func (p *PersonThinBuilder) BuildLeg() {
	fmt.Println("瘦子的腿")
}
func (p *PersonThinBuilder) BuildBody() {
	fmt.Println("瘦子的身子")
}

type PersonDirector struct {
	pb PersonBuilder
}

func (a *PersonDirector) NewPersonDirector(pb PersonBuilder) {
	a.pb = pb
}
func (a *PersonDirector) CreatePerson() {
	a.pb.BuildHead()
	a.pb.BuildBody()
	a.pb.BuildArm()
	a.pb.BuildLeg()
}
