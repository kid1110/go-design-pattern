package proxy

import "testing"

func Test1(t *testing.T) {
	girl := Girl{Name: "张三"}
	k := NewProxy(girl)
	k.GiveChocolate()
	k.GiveDolls()
	k.GiveFlowers()
}
