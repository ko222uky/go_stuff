package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("Hello")
}


func main() {
	go sayHello()
	runtime.Gosched()	// This calls the Go scheduler, which executes when triggered by user-level events
				// This is separate from the OS scheduler. Here, we manually call the Go scheduler.
				// The main() function in this scenario would terminate before giving the goroutine a chance to execute.
				// By calling the goscheduler, we increase the chances (but do not guarantee) the execution of our goroutine.
				// NOTE: the scheduler may simply choose the main() goroutine again, and we won't see "Hello" print out.
	fmt.Println("Finished")

}
