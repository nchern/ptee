package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var outputFile = flag.String("o", "", "filename to redirect output of a given command to")

func init() {
	flag.Usage = func() {
		fmt.Println("Starts a process(the 1st argument specifies the command to run),")
		fmt.Println("connects stdin to this process' stdin and mirrors it to stdout as well.")
		fmt.Println()
		fmt.Printf("Usage: %s [command] [args...]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(2)
	}
}

func run() error {
	in := io.TeeReader(os.Stdin, os.Stdout)
	cmd := exec.Command(flag.Args()[0], flag.Args()[1:]...)
	cmd.Stdin = in
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			return err
		}
		defer f.Close()
		cmd.Stdout = f
	}
	return cmd.Run()
}

func main() {
	if err := run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
