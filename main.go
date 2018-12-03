
package main

import (
"fmt"    
"os/exec"
"bytes"
"encoding/json"
)
type AddressGenJson struct {
	Meta struct {
		Coin       string `json:"coin"`
		CryptoType string `json:"cryptoType"`
		Encrypted  string `json:"encrypted"`
		Filename   string `json:"filename"`
		Label      string `json:"label"`
		LastSeed   string `json:"lastSeed"`
		Secrets    string `json:"secrets"`
		Seed       string `json:"seed"`
		Tm         string `json:"tm"`
		Type       string `json:"type"`
		Version    string `json:"version"`
	} `json:"meta"`
	Entries []struct {
		Address   string `json:"address"`
		PublicKey string `json:"public_key"`
		SecretKey string `json:"secret_key"`
	} `json:"entries"`
}

func main() {
	cmd := exec.Command("skycoin-cli", "addressGen")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	//	Run the command
	cmd.Run()

	//	Output our results
	

	jsonOut := out.String()
	var all AddressGenJson
	fmt.Printf("JsonOut: %v", jsonOut) 	
	json.Unmarshal([]byte(jsonOut), &all)
	fmt.Println(all)
	
	s:= &all	
	seed := s.Meta.Seed
	address:= s.Entries
	fmt.Println(seed)
	fmt.Println(address)

	

}
