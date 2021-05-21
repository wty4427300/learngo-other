package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- requestVote1()
		}()
	}
	for i := 0; i < 10; i++ {
		v := <-ch
		if v {
			count += 1
		}
	}
	if count >= 5 {
		println("received 5+ votes!")
	} else {
		println("lost")
	}
}

func requestVote1() bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return rand.Int()%2 == 0
}
