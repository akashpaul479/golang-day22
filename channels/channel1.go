package channels

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	Value string
	Err   error
}

func worker3(url string, wg *sync.WaitGroup, resultchan chan result) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed: %s\n", url)

	resultchan <- result{
		Value: url,
		Err:   nil,
	}
}

func Fanin2() {

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
	resultchan := make(chan result, 50)

	startTime := time.Now()

	for _, job := range jobs {
		wg.Add(1)
		go worker3(job, &wg, resultchan)
	}

	wg.Wait()
	close(resultchan)

	for result := range resultchan {
		fmt.Printf("received: %s\n", result)
	}

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
