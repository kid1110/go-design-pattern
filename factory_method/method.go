package factoryMethod

//此处代码与简单工厂模式做对比，所以也用计算器作为例子

type Operation interface {
	GetResult(a, b float64) float64
}
type Factory interface {
	CreateOperation()
}

type Sub struct{}
type Min struct{}

type AddFactory struct {
}
type MinFactory struct {
}

func (a *AddFactory) CreateOperation() Sub {
	return Sub{}
}
func (m *MinFactory) CreateOperation() Min {
	return Min{}
}
func (*Sub) GetResult(a, b float64) float64 {
	return a + b
}

func (*Min) GetResult(a, b float64) float64 {
	return a - b
}
