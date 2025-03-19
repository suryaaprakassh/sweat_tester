package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner interface {
	Exec(string, string) error
}

type CppRunner struct {
	filePath string
}

func NewCppRunner() *CppRunner {
	return &CppRunner{
		filePath: "./codes/code.cpp",
	}
}

func (c *CppRunner) Compile() error {
	cmd := exec.Command("clang++", c.filePath, "-o", "main")
	_, err := cmd.CombinedOutput()
	return err
}

func (c *CppRunner) Exec(input string, output string) error {

	if err := c.Compile(); err != nil {
		return err
	}
	cmd := exec.Command("./main")
	cmd.Stdin = strings.NewReader(input)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	outLines := strings.Split(output, "\n")
	cmdLines := strings.Split(output, "\n")
	if len(outLines) != len(cmdLines) {
		return fmt.Errorf("cpp: want %d output, got %d output", len(outLines), len(cmdLines))
	}

	for i := range len(outLines) {
		if outLines[i] != cmdLines[i] {
			return fmt.Errorf("cpp: output did not match line:%d, want %s , got: %s", i+1, outLines[i], cmdLines[i])
		}
	}

	CreateFile("cpp_out", string(cmdOutput))
	return nil
}
