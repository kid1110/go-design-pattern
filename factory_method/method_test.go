package factoryMethod

import (
	"fmt"
	"testing"
)

func TestSub(t *testing.T) {
	operFactory := AddFactory{}
	oper := operFactory.CreateOperation()
	data := oper.GetResult(1, 2)
	fmt.Println(data)
}

func TestMin(t *testing.T) {
	operFactory := MinFactory{}
	oper := operFactory.CreateOperation()
	data := oper.GetResult(1, 2)
	fmt.Println(data)
}
