package main

/*
Construct a pipeline that connects an arbitrary number of
goroutines with channels. What is the maximum number of pipeline stages you can create without
running out of memory? How long does a value take to transit the entire pipeline?
*/
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	head := make(chan time.Time)
	last := head
	f, err := os.Create("pipeline_routines.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	for stageCount := 1; ; stageCount++ {
		go pipeLine(last, stageCount, f)
		head <- time.Now()

		temp := last
		last = make(chan time.Time)
		go connectPipes(temp, last)
	}
}

func pipeLine(last chan time.Time, stageCount int, writer *os.File) {
	startTime := <-last
	endTime := time.Now()
	fmt.Printf("GoRoutine #: %d\t Time: %v\n", stageCount, endTime.Sub(startTime))
	writer.WriteString(strconv.Itoa(stageCount) + "," + fmt.Sprint(endTime.Sub(startTime)) + "\n")
}

func connectPipes(src, dst chan time.Time) {
	for t := range src {
		dst <- t
	}
}
