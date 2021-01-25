package main

import (
	"io"
	"os"
	"os/exec"
)

func execution(commandName string, args []string, stdout, stderr io.Writer) (*os.ProcessState, error) {
	cmd := exec.Command(commandName, args...)
	childStdout, _ := cmd.StdoutPipe()
	childStderr, _ := cmd.StderrPipe()

	go io.Copy(stdout, childStdout)
	go io.Copy(stderr, childStderr)

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return cmd.ProcessState, nil
}
