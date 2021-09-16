# BSDISYS_Philosophers
Go implementation of "The Dining Philosophers".  
Implemented by "Vildt fed gruppe".
## Instructions
Run program by running `go run .` in the terminal in the location.  
To enable or disable debug 'eating' messages, change the `ANNOUNCEEATING` boolean variable in `Constants.go`
## Implementation
For avoiding deadlocks, when philosophers decide to eat, they attempt to begin eating, but if they cannot get a hold of two forks, they return to thinking. When a philosopher is attempting to eat, only they can interact with the forks on the table. Other philosopher will have to wait for that philosopher to begin eating or return to thinking.  
A thread is created for each philosopher, which is constantly rolling a virtual die. When it rolls correctly, it'll make it's associated philosopher attempt to begin eating, or lay down their forks.
