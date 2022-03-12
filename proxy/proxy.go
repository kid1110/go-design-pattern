package proxy

import "fmt"

type Gift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}
type Girl struct {
	Name string
}

type Pursuit struct {
	Gift
	SchoolGirl Girl
}
type Proxy struct {
	Gift
	Pu Pursuit
}

func (p *Pursuit) GiveDolls() {
	fmt.Println(p.SchoolGirl.Name + "give dolls")
}

func (p *Pursuit) GiveFlowers() {
	fmt.Println(p.SchoolGirl.Name + "give flowers")
}

func (p *Pursuit) GiveChocolate() {
	fmt.Println(p.SchoolGirl.Name + "give chocolate")
}

func (pr *Proxy) GiveDolls() {
	pr.Pu.GiveDolls()
}
func (pr *Proxy) GiveFlowers() {
	pr.Pu.GiveFlowers()
}
func (pr *Proxy) GiveChocolate() {
	pr.Pu.GiveChocolate()
}
func NewProxy(girl Girl) Proxy {
	pu := Pursuit{SchoolGirl: girl}
	return Proxy{
		Pu: pu,
	}
}
