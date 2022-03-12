package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

func Test1(t *testing.T) {
	one := GetInstance()
	two := GetInstance()
	if one != two {
		t.Fatal("instance is not equal")
	}
}

func TestParrallel(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal(("instance is not equal"))
		}
	}
}
