package main

import (
	"fmt"
	"log/slog"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type JavaRunner struct {
	filePath  string
	className string
}

func NewJavaRunner() *JavaRunner {
	return &JavaRunner{
		filePath:  "./Main.java",
		className: "Main",
	}
}

func (j *JavaRunner) Change() {
	cmd := exec.Command("cp", "code.java", "Main.java")
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(err.Error())
		slog.Info(string(out))
	}
}

func (j *JavaRunner) Unchange() {
	cmd := exec.Command("rm", "Main.java")
	_, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error(err.Error())
	}
}

func (j *JavaRunner) Compile() error {
	cmd := exec.Command("javac", j.filePath)
	_, err := cmd.CombinedOutput()
	return err
}

func (j *JavaRunner) Exec(input string, output string) error {
	j.Change()
	if err := j.Compile(); err != nil {
		return fmt.Errorf("compilation error: %w", err)
	}
	defer j.Unchange()

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
			log.Printf("Java: output did not match line:%d, want %s , got: %s", i+1, outLines[i], cmdLines[i])
		}
	}

	CreateFile("java_out", string(cmdOutput))
	return nil
}
