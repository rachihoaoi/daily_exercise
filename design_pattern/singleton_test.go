package design_pattern

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Singleton struct {
	Time time.Time
}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton{
	once.Do(func() {
		singleton = &Singleton{
			Time: time.Now(),
		}
	})

	return singleton
}


func TestSingleton(t *testing.T) {
	ins1 := GetInstance().Time
	fmt.Println(ins1)
	time.Sleep(time.Second)
	ins2 := GetInstance().Time
	fmt.Println(&ins2)
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}