package decorator

import "fmt"

type Component interface {
	Show()
}
type Person struct {
	Component
	Name string
}
type Finery struct {
	Component
}
type BigTrouser struct {
	Finery
}
type TShirts struct {
	Finery
}

func (f *Finery) Decorate(p Component) {
	f.Component = p
}
func (f *Finery) Show() {
	f.Component.Show()
}
func (p *Person) Show() {
	fmt.Println("装扮的", p.Name)
}
func (b *BigTrouser) Show() {
	fmt.Println("大裤衩子")
	b.Finery.Show()

}
func (t *TShirts) Show() {
	fmt.Println("大tshirts")
	t.Finery.Show()

}
