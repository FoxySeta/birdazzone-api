package main

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func TestMain(m *testing.M) {
	//util.TestWithServer(m)
	sttyArgs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil,
	}

	pid, err := syscall.ForkExec("/bin/go", []string{"/bin/go", "run", ""}, &sttyArgs)
	fmt.Println(pid)
	fmt.Println(err)

	os.Exit(m.Run())
}
