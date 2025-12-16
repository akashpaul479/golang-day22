package workerpool

import (
	"fmt"
	"sync"
	"time"
)

type result1 struct {
	Value string
	Err   error
}

func worker4(jobschan chan string, wg *sync.WaitGroup, resultchan chan result) {
	defer wg.Done()

	for job := range jobschan {
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("image processed: %s\n", job)
		resultchan <- result{
			Value: job,
			Err:   nil,
		}
	}
	fmt.Printf("Worker shutting down")
}

func Fanin4() {

	jobs := []string{
		"imag_1.png",
		"imag_2.png",
		"imag_3.png",
		"imag_4.png",
		"imag_5.png",
		"imag_6.png",
		"imag_7.png",
		"imag_8.png",
		"imag_9.png",
		"imag_10.png",
		"imag_11.png",
		"imag_12.png",
		"imag_13.png",
		"imag_14.png",
		"imag_15.png",
	}
	var wg sync.WaitGroup
	totalworkers := 5

	resultchan := make(chan result, 50)
	jobschan := make(chan string, len(jobs))

	startTime := time.Now()

	for i := 1; i <= totalworkers; i++ {
		wg.Add(1)
		go worker4(jobschan, &wg, resultchan)
	}
	go func() {
		wg.Wait()
		close(resultchan)
	}()

	for i := 0; i < len(jobs); i++ {
		jobschan <- jobs[i]

	}
	close(jobschan)

	for result := range resultchan {
		fmt.Printf("job completed %v\n", result)
	}

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
