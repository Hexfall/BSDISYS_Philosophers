package main

import (
	"fmt"
	"math/rand"
)

var philEating [PHILCOUNT]bool
var philTimes  [PHILCOUNT]int
var philIn [PHILCOUNT]chan int
var philOut [PHILCOUNT]chan int

// Returns whether a philosopher is currently eating, and how often he's eaten.
func getPhilInfo(id *int) (bool, int) {
	philIn[*id] <- 2
	eating := <- philOut[*id] == 1
	philIn[*id] <- 3
	eaten := <- philOut[*id]
	return eating, eaten
}

// Sets up channels and start goroutines for philosophers.
func startPhilGoroutines() {
	for i := 0; i < PHILCOUNT; i++ {
		philIn[i]  = make(chan int)
		philOut[i] = make(chan int)
		go philCommGoroutine(i)
		go philosopherPonderanceGoroutine(i)
	}
}

// Takes input through philIn[id] and does as follows:
// 0: Make philosopher attempt to eat.
// 1: Make philosopher stop eating.
// 3: Query philosopher whether he's currently eating or not.
// 4: Query philosopher how often he's eaten.
// Responds through philOut[id], and always sends a response after input.
func philCommGoroutine(id int) {
	var command, response int
	for {
		command = <- philIn[id]
		switch command {
		case 0: // Attempt to begin eating.
			success := tryEat(id)
			if success {
				philEating[id] = true
				philTimes[id]++
				if ANNOUNCEEATING {
					fmt.Printf("Philosopher %d started eating\n", id)
				}
			}
		case 1: // Stops eating and returns to deliberating.
			stopEating(id)
			philEating[id] = false
			if ANNOUNCEEATING {
				fmt.Printf("Philosopher %d stopped eating\n", id)
			}
		case 2: // Query philosopher whether he's eating or not.
			response = btoi(philEating[id])
		case 3: // Query philosopher how often he has eaten so far.
			response = philTimes[id]
		}
		philOut[id] <- response
	}
}

// Attempts to imitate the whimsy of a philosopher, switching between eating and speculating seemingly at random.
func philosopherPonderanceGoroutine(id int) {
	for {
		if rand.Float64() < PHILSWITCHCHANCE {
			philIn[id] <- 2
			isEating := <- philOut[id] == 1
			// Switch: Thinking <-> Eating.
			if isEating {
				// Drop forks and return to positing on the nature of the universe.
				philIn[id] <- 1
			} else {
				// Attempt to begin eating. Return to postulating, if missing fork.
				philIn[id] <- 0
			}
			<- philOut[id]
		}
	}
}
