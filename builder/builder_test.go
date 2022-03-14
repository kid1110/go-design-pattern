package builder

import "testing"

func Test1(t *testing.T) {

	a := PersonThinBuilder{}
	b := PersonDirector{}
	b.NewPersonDirector(&a)
	b.CreatePerson()
}
