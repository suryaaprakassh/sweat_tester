package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type JavaRunner struct {
	filePath string
	className string
}

func NewJavaRunner() *JavaRunner {
	return &JavaRunner{
		filePath: "./codes/Main.java",
		className: "Main", 
	}
}

func (j *JavaRunner) Compile() error {
	cmd := exec.Command("javac", j.filePath)
	_, err := cmd.CombinedOutput()
	return err
}

func (j *JavaRunner) Exec(input string, output string) error {
	if err := j.Compile(); err != nil {
		return fmt.Errorf("compilation error: %w", err)
	}

	dir := filepath.Dir(j.filePath)
	cmd := exec.Command("java", "-cp", dir, j.className)
	cmd.Stdin = strings.NewReader(input)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("execution error: %w", err)
	}
	
	outLines := strings.Split(output, "\n")
	cmdLines := strings.Split(string(cmdOutput), "\n")

	if len(outLines) != len(cmdLines) {
		return fmt.Errorf("Java: want %d output lines, got %d output lines", len(outLines), len(cmdLines))
	}
	
	for i := range outLines {
		if outLines[i] != cmdLines[i] {
			return fmt.Errorf("Java: output did not match line:%d, want %s , got: %s", i+1, outLines[i], cmdLines[i])
		}
	}
	
	CreateFile("java_out", string(cmdOutput))
	return nil
}
