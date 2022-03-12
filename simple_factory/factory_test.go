package simplefactory

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	data := NewOperation("+")
	res := data.GetResult(1, 2)
	fmt.Println(res)
}

func Test2(t *testing.T) {
	data := NewOperation("-")
	res := data.GetResult(1, 2)
	fmt.Println(res)

}
