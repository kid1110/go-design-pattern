package singleton

import "sync"

type Singleton interface {
	say()
}
type singleton struct {
}

func (s singleton) say() {

}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
