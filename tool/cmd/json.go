package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func printJson(v interface{}) {
	result, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(result))
}
