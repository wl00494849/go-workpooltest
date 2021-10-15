package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, job <-chan int, wg sync.WaitGroup) {
	for j := range job {
		fmt.Println("worker", id, " start job", j)
		time.Sleep(time.Second * 3)
		fmt.Println("worker", id, " finish job", j)
		wg.Done()
	}
}

func main() {
	jobCount := 20
	jobs := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(jobCount)

	//start 3 worker
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, wg)
	}

	//send job to channl
	for i := 1; i < jobCount; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}
