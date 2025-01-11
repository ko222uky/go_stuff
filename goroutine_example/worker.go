package main
import ( 
	"fmt"
	"time"
)

/*
From "Learn Concurrent Programming with Go," page 30:

Simulates computational work via sleeping for 1 second.

*/
func doWork(id int) {

	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Work %d completed at %s\n", id, time.Now().Format("15:04:05"))

}

// Simulate sequential processes...
func main() {
	fmt.Printf("Synchronous calls\n") // expect completion in ~ 5 seconds   
	for i := 0; i < 5; i++ {
		doWork(i)
	}

	fmt.Printf("Asynchronous calls\n") // process does not wait for function return to move to next instruction
	// expect completion in ~ 2 seconds
	for i := 0; i < 5; i++ {
		go doWork(i) // starts new goroutine that calls our doWork() function
	}
	time.Sleep(2 * time.Second) // We need this because if the main execution finishes, the goroutines won't get to run.
}

