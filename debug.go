package utp

import (
	"fmt"
	"runtime/debug"
	"time"
)

type watchDog struct {
	*time.Timer
}

func newWatchDog(d time.Duration) *watchDog {
	s := append([]byte(nil), debug.Stack()...)
	return &watchDog{Timer: time.AfterFunc(d, func() {
		fmt.Println(string(s))
		panic("bowwow")
	})}
}
