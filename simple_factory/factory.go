package simplefactory

import (
	"fmt"
)

//两数加减运算

type Operation interface {
	GetResult(a, b float64) float64
}

type Sub struct{}
type Min struct{}

func (*Sub) GetResult(a, b float64) float64 {
	return a + b
}

func (*Min) GetResult(a, b float64) float64 {
	return a - b
}

func NewOperation(method string) Operation {
	if method == "+" {
		return &Sub{}
	} else if method == "-" {
		return &Min{}
	} else {
		fmt.Println("暂无该选项")
		return &Sub{}
	}
}
