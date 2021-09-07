package main

import (
	"fmt"
	"sync"
)

var forkLocks [PHILCOUNT]sync.Mutex

func use_fork() {
	fmt.Println("no")
}
