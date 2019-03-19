package main

import (
	"fmt"

	"github.com/jeromer/syslogparser/rfc3164"
)

func main() {
	b := `<34>Mar 17 10:19:03 localhost sshd[16155]: Failed password for root from 82.165.30.36 port 53644 ssh2`
	buff := []byte(b)

	p := rfc3164.NewParser(buff)
	err := p.Parse()
	if err != nil {
		panic(err)
	}

	for k, v := range p.Dump() {
		fmt.Println(k, ":", v)
	}
}
