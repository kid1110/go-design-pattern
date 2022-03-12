package decorator

import (
	"testing"
)

func Test1(t *testing.T) {
	xc := &Person{Name: "张三"}
	kk := &BigTrouser{}
	d := &TShirts{}

	kk.Decorate(xc)
	d.Decorate(kk)
	d.Show()

}
