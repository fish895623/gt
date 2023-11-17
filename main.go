package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Commit() {
	output := exec.Command("git", "commit")
	output.Stdin = os.Stdin
	output.Stdout = os.Stdout
	output.Stderr = os.Stderr
	if err := output.Start(); err != nil {
		fmt.Println(err)
	}
	if err := output.Wait(); err != nil {
		fmt.Println(err)
	}
	if stdout, err := output.Output(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(string(stdout))
	}
}

func main() {
	exec.Command("git", "fetch", "--all").Run()
	Commit()
	exec.Command("git", "push").Run()
}
