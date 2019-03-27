package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgSent() string {
	var msg [5]string
	msg[0] = "this is msg 0"
	msg[1] = "this is msg 1"
	msg[2] = "this is msg 2"
	msg[3] = "this is msg 3"
	msg[4] = "this is msg 4"
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return msg[r1.Intn(5)]
}

func main() {
	for {
		time.Sleep(time.Millisecond * 200)
		fmt.Println(msgSent())
	}
}
