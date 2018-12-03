package main

import (
"fmt"    
"os/exec"
"bytes"
)

func main() {
	cmd := exec.Command("skycoin-cli", "addressGen")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	//	Run the command
	cmd.Run()

	//	Output our results
	fmt.Printf("Result: %v / %v", out.String(), stderr.String())
}
