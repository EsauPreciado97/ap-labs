/*
 Write a program with two goroutines that send messages back and forth over
two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var pings string

	ping := make(chan int)
	pong := make(chan int)
	done := make(chan struct{})

	flag.Parse()
	pings = flag.Arg(0)

	pingpongCount, err := strconv.Atoi(pings)
	if err != nil {
		// handle error
		fmt.Println("Run As: go run ping-pong.go <number of pings>")
		os.Exit(2)
	}

	startTime := time.Now()

	go func() {
		for n := 0; n < pingpongCount; n++ {
			ping <- n
			//fmt.Printf("Ping! ")
			<-pong
		}
		close(ping)
		close(done)
	}()

	go func() {
		for n := range ping {
			//fmt.Printf("Pong!\n")
			pong <- n
		}

		close(pong)
	}()

	<-done

	endTime := time.Now()
	deltaT := endTime.Sub(startTime)

	time := float64(deltaT.Nanoseconds()) / 1000000000.0
	rate := float64(pingpongCount) / time

	fmt.Printf("Time: %v \t Message number: %v Replies: \t%f \n", deltaT, pingpongCount, rate)
}
