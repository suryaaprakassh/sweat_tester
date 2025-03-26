package main

import (
	"fmt"
	"log"
	"log/slog"
	"os/exec"
	"strings"
)

type PythonRunner struct {
	filePath string
}

func NewPythonRunner() *PythonRunner {
	return &PythonRunner{
		filePath: "./code.py",
	}
}

func (j *PythonRunner) Exec(input string, output string) error {

	cmd := exec.Command("python3", j.filePath)
	cmd.Stdin = strings.NewReader(input)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Got", "out", cmdOutput)
		return fmt.Errorf("execution error: %w", err)
	}

	outLines := strings.Split(output, "\n")
	cmdLines := strings.Split(string(cmdOutput), "\n")

	if len(outLines) != len(cmdLines) {
		return fmt.Errorf("Python: want %d output lines, got %d output lines", len(outLines), len(cmdLines))
	}

	for i := range outLines {
		if outLines[i] != cmdLines[i] {
			log.Printf("Python: output did not match line:%d, want %s , got: %s", i+1, outLines[i], cmdLines[i])
		}
	}

	CreateFile("py_out", string(cmdOutput))
	return nil
}
