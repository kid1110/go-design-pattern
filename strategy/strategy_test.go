package strategy

import "testing"

func Test1(t *testing.T) {
	data := PayContext{&CashNormal{}}
	data.Data.PayAlgorithm()
}

func Test2(t *testing.T) {
	data := PayContext{&CashRebate{}}
	data.Data.PayAlgorithm()
}
