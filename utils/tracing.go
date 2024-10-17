package utils

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

/*
This package allows you to trace part of the code and view the 
execution path then the code run completes

To trace eg a function add at the top of the function 
	f := utils.StartTrace()
	defer utils.StopTrace(f)



When the program completes a trace.out file will be created
To view this file run the following in the command line
 go tool trace trace.out
*/

// starts the execution trace
func StartTrace() *os.File {
	// Create a file to store the trace output
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}

	// Start the trace
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	fmt.Println("Tracing started")
	return f
}

// closes the trace File and stops the execution trace
func StopTrace(f *os.File) {
	trace.Stop()
	f.Close()
	fmt.Println("trace cleanup complete")
}
