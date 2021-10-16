package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func init() {
	flag.Usage = func() {
		fmt.Println("Starts a process(the 1st argument specifies the command to run),")
		fmt.Println("connects stdin to this process' stdin and mirrors it to stdout as well.")
		fmt.Println()
		fmt.Printf("Usage: %s [command] [args...]\n", os.Args[0])
	}
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
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
