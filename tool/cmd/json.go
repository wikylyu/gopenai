package cmd

import (
	"encoding/json"
	"fmt"
)

func printJson(v interface{}) {
	result, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(result))
}
