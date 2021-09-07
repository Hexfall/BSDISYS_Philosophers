package main

import (
	"fmt"
	"time"
)

func main() {
	startGoroutines()
	interactive()
}

func startGoroutines() {
	startForkGoroutines()
	startPhilGoroutines()
}

func forkMap() {
	for {
		fmt.Print("\r")
		for i := 0; i < PHILCOUNT; i++{
			var app = 'T'
			if ForkBusy[i] {
				app = 'E'
			}
			fmt.Printf("%c ", app)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

// Interactive mode for console. Can query for current program status.
func interactive() {
	var com string
	var id int
	fmt.Printf("%d philosophers are pondering and eating away in your kitchen.\n", PHILCOUNT)
	fmt.Println("Write \"phil/fork {id}\" to get information about that philosopher/fork, or write \"q\" to kick the philosophers out of your house.")
	for {
		fmt.Scanf("%s %d", &com, &id)
		id = limit(id)
		switch com {
		case "phil":
			eating, eaten := getPhilInfo(&id)
			if eating {
				fmt.Printf("Philosopher %d is currently eating.\n", id)
			} else {
				fmt.Printf("Philosopher %d is currently not eating.\n", id)
			}
			fmt.Printf("Philosopther %d has eaten %d times so far\n", id, eaten)
		case "fork":
			inUse, timesUsed := getForkInfo(&id)
			if inUse {
				fmt.Printf("Fork %d is currently is use.\n", id)
			} else {
				fmt.Printf("Fork %d is currently not in use.\n", id)
			}
			fmt.Printf("Fork %d has been used %d times so far.\n", id, timesUsed)
		case "q":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

// Use "go run ." in console to run