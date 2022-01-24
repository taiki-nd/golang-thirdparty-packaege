package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	//同時に走るgoroutineを制御する
	/*
		if err := s.Acquire(ctx, 1); err != nil {
			fmt.Println(err)
			return
		}
	*/
	//goroutineの本数をを制御する
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("could not get lock")
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("DONE")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
