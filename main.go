package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [command] [args...]\n", os.Args[0])
		os.Exit(2)
	}
}

func main() {
	in := io.TeeReader(os.Stdin, os.Stdout)
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
