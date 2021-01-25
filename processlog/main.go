package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	logDir, CommandName, args := parseArgs()
	stdout, stderr, err := initOut(logDir, CommandName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create file: %v\n", err)
		os.Exit(1)
	}

	startTime := time.Now()
	ps, err := execution(CommandName, args, stdout, stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't exec command: %s\n", CommandName)
		os.Exit(1)
	}

	exitTime := time.Now()
	fmt.Fprintln(stdout, ps.String())
	fmt.Fprintf(stdout, "wall clock time=%v system time=%v user time=%v\n",
		exitTime.Sub(startTime), ps.SystemTime(), ps.UserTime())
}
