package main

import (
	"bytes"
	"fmt"
	"scope/internal"

	"gopkg.in/yaml.v3"
)

func main() {
	report, err := internal.BuildReport()
	if err != nil {
		fmt.Printf("Error building report: %+v", err)
		return
	}

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)

	err = encoder.Encode(report)
	if err != nil {
		fmt.Printf("Error marshaling to YAML: %+v", err)
		return
	}

	fmt.Print(buf.String())
}
