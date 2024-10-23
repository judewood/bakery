package tracing

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/trace"
)

// StartTrace starts the execution trace
func StartTrace() *os.File {
	// Create a file to store the trace output
	f, err := os.Create("trace.out")
	if err != nil {
		slog.Warn("failed to create trace output file: %v", "error", err)
	}

	// Start the trace
	if err := trace.Start(f); err != nil {
		slog.Warn("failed to start trace: %v", "error", err)
	}
	fmt.Println("Tracing started")
	return f
}

// StopTrace closes the trace File and stops the execution trace
func StopTrace(f *os.File) {
	trace.Stop()
	f.Close()
	fmt.Println("trace cleanup complete")
}
