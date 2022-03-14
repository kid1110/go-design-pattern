package observer

import "testing"

func Test1(t *testing.T) {

	a := NewConcreteSubject()
	b := ConcreteObserver{}
	a.Attach(&b)
	a.Notify()
}
