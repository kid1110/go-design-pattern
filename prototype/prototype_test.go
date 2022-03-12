package prototype

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	a := Resume{Name: "kid", Sex: "boy"}
	b := a.Clone()
	fmt.Println(a)
	fmt.Println(b)
	b.Name = "zzzzz"
	fmt.Println(b)
}
