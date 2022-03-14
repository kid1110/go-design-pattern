package observer

import "fmt"

type Subject interface {
	Notify()
	Attach()
}

type Observer interface {
	Update()
}
type ConcreteObserver struct {
	Observer
}

type ConcreteSubject struct {
	Subject
	observers []Observer
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{
		observers: make([]Observer, 0),
	}
}
func (c *ConcreteSubject) Notify() {
	for _, o := range c.observers {
		o.Update()
	}
}
func (c *ConcreteObserver) Update() {
	fmt.Println("观察者1号")
}
func (c *ConcreteSubject) Attach(o Observer) {
	c.observers = append(c.observers, o)
}
