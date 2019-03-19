package main

import (
	"fmt"

	"github.com/google/gonids"
)

func main() {
	rule := `alert tcp $HOME_NET any -> $EXTERNAL_NET 80 (msg:"GONIDS TEST hello world"; flow:established,to_server; content:"hello world"; classtype:trojan-activity; sid:1; rev:1;)`

	r, err := gonids.ParseRule(rule)
	if err != nil {
		// Handle parse error
	}
	// Do something with your rule.
	switch r.Action {
	case "alert":
		fmt.Println("alert")
		fmt.Println(r.SID)
	case "drop":
		// This is a 'drop' rule.
	case "pass":
		// This is a 'pass' rule.
	default:
		// I have no idea what this would be. =)
	}
}
