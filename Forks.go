package main

import (
	"sync"
)

var forkUses [PHILCOUNT]int
var ForkBusy [PHILCOUNT]bool
var forkMutex sync.Mutex

// Returns whether a fork is currently in use, and how often it's been used.
func getForkInfo(id *int) (bool, int) {
	forkIn[*id] <- 2
	inUse := <- forkOut[*id] == 0
	forkIn[*id] <- 3
	used := <- forkOut[*id]
	return inUse, used
}

// Channels for communicating with fork coroutines. Always responds after input.
var forkIn  = [PHILCOUNT]chan int{}
var forkOut = [PHILCOUNT]chan int{}
// Start goroutines for forks.
func startForkGoroutines() {
	for i, _ := range forkIn {
		forkIn[i]  = make(chan int)
		forkOut[i] = make(chan int)
		go forkGoroutine(i)
	}
}

// A philosopher attempts to begin eating. Does not claim any forks, unless he can claim both.
// Returns true if he can (and does) begin eating.
func tryEat(philId int) bool {
	forkMutex.Lock()
	var ret bool

	if canEat(philId) {
		useForks(philId)
		ret = true
	}
	forkMutex.Unlock()
	return ret
}

// Returns true if the forks adjacent to given philosopher are available.
// WARNING: Does not lock variable availability and should be considered unreliable outside tryEat.
func canEat(philId int) bool {
	var forkLeft, forkRight bool
	forkIn[limit(philId)]     <- 2
	forkIn[limit(philId + 1)] <- 2
	forkLeft  = <- forkOut[limit(philId)]     == 1
	forkRight = <- forkOut[limit(philId + 1)] == 1
	return forkLeft && forkRight
}

// Use the forks adjacent to given philosopher.
// WARNING: Should only be called from tryEat. If you want your philosopher to eat, call tryEat.
func useForks(philId int) {
	forkIn[limit(philId)]     <- 1
	forkIn[limit(philId + 1)] <- 1
	<- forkOut[limit(philId)]
	<- forkOut[limit(philId + 1)]
}

// Makes a philosopher lay down his forks.
func stopEating(philId int) {
	forkIn[limit(philId)]     <- 0
	forkIn[limit(philId + 1)] <- 0
	<- forkOut[limit(philId)]
	<- forkOut[limit(philId + 1)]
}

// Takes input through forkIn[id] and does as follows:
// 0: Make fork available.
// 1: Make fork unavailable.
// 2: Query fork, whether it's available. (1 for yes, 0 for no)
// 3: Query fork how many times it's been used.
// Responds through forkOut[id], and always sends a response after input.
func forkGoroutine(id int) {
	var command, response int
	for {
		response = 0
		command = <- forkIn[id]
		switch command {
		case 0: // Make available.
			if !ForkBusy[id] {
				// Fork is already idle, throw error.
			}
			ForkBusy[id] = false
		case 1: // Make unavailable.
			if ForkBusy[id] {
				// Fork is already busy, throw error.
			}
			ForkBusy[id] = true
			forkUses[id]++
		case 2: // Get availability.
			response = btoi(!ForkBusy[id])
		case 3: // Get usage.
			response = forkUses[id]
		}
		forkOut[id] <- response
	}
}