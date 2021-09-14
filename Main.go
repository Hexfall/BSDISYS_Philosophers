package main

import (
	"fmt"
	"time"
)

var runFeed bool

func main() {
	startGoroutines()
	interactive()
}

func startGoroutines() {
	startForkGoroutines()
	startPhilGoroutines()
	go liveFeed()
}

func liveFeed() {
	var chars [PHILCOUNT * 2 + 1]rune
	for {
		time.Sleep(50 * time.Millisecond)
		if !runFeed {
			continue
		}
		for i := 0; i < PHILCOUNT; i++ {
			forkIn[i] <- 2
			if 0 == <- forkOut[i] {
				chars[i*2] = 'F'
			} else {
				chars[i*2] = 'f'
			}
			philIn[i] <- 2
			if 1 == <- philOut[i] {
				chars[i*2 + 1] = 'P'
			} else {
				chars[i*2 + 1] = 'p'
			}
		}
		chars[PHILCOUNT * 2] = chars[0]
		fmt.Print("\r")
		for _, c := range chars {
			fmt.Printf("%c", c)
		}
	}
}

// Interactive mode for console. Can query for current program status.
func interactive() {
	var com string
	var id int
	fmt.Printf("%d philosophers are pondering and eating away in your kitchen.\n", PHILCOUNT)
	fmt.Println("Write \"phil/fork {id}\" to get information about that philosopher/fork, or write \"q\" to kick the philosophers out of your house.")
	fmt.Println("You can also type \"live\" to get a live feed of the philosophers and their forks.")
	for {
		fmt.Scan(&com)
		switch com {
		case "phil":
			fmt.Scanf("%d", &id)
			id = limit(id)
			eating, eaten := getPhilInfo(&id)
			if eating {
				fmt.Printf("Philosopher %d is currently eating.\n", id)
			} else {
				fmt.Printf("Philosopher %d is currently not eating.\n", id)
			}
			fmt.Printf("Philosopther %d has eaten %d times so far\n", id, eaten)
		case "fork":
			fmt.Scanf("%d", &id)
			id = limit(id)
			inUse, timesUsed := getForkInfo(&id)
			if inUse {
				fmt.Printf("Fork %d is currently is use.\n", id)
			} else {
				fmt.Printf("Fork %d is currently not in use.\n", id)
			}
			fmt.Printf("Fork %d has been used %d times so far.\n", id, timesUsed)
		case "live":
			runFeed = true
			fmt.Scanln()
			runFeed = false
		case "q":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

// Use "go run ." in console to run