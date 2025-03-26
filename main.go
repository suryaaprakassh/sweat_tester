package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func CreateFile(name string, content string) error {
	return os.WriteFile(name, []byte(content),0644)
} 

type Result struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type ResFile []Result

func (r *ResFile) Read(fileName string) error {
	data, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, r)
	if err != nil {
		return err
	}

	return nil
}

func (r ResFile) Debug() {
	for _, val := range r {
		fmt.Printf("%+v\n", val)
	}
}

func (r ResFile) Exec(runner Runner) {
	err:=runner.Exec(r.GetInput(),r.GetOutput())
	if err != nil {
		slog.Error(err.Error())
	}
}

func (r ResFile) GetInput() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%d\n", len(r))
	for _, val := range r {
		fmt.Fprintf(&b, "%s\n", val.Input)
	}
	return b.String()
}

func (r ResFile) GetOutput() string {
	var b strings.Builder
	for _, val := range r {
		fmt.Fprintf(&b, "%s\n", val.Output)
	}

	return b.String()
}

func main() {
	var file ResFile
	if err := file.Read("results.json"); err != nil {
		slog.Error(err.Error())
	}
	CreateFile("inp", file.GetInput())
	CreateFile("out",file.GetOutput())

	
	rn:=NewCppRunner()
	jrn:=NewJavaRunner()
	prn:=NewPythonRunner()

	file.Exec(rn)
	file.Exec(jrn)
	file.Exec(prn)
}
