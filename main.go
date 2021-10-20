package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var n nix
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error decoding nix: %v", err)
		os.Exit(1)
	}
	c, err := NewCycloneFromNix(&n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error converting nix to cyclone: %v", err)
		os.Exit(1)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(c)
	return
}
