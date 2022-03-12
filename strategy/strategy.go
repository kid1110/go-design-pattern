package strategy

import "fmt"

//正常收费，打折，等问题

//Strategy定义所有支持算法的公共接口
type Strategy interface {
	PayAlgorithm()
}

//具体实现
type CashNormal struct{}

//具体实现
type CashRebate struct{}

type PayContext struct {
	Data Strategy
}

func (p *CashNormal) PayAlgorithm() {
	fmt.Println("正常返还")
}
func (ctx *PayContext) ContextInterface(pay Strategy) {
	pay.PayAlgorithm()
}

func (p *CashRebate) PayAlgorithm() {
	fmt.Println("打折返还")
}
