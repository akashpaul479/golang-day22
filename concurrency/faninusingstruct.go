package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type result struct {
	Value string
	Err   error
}

func worker2(url string, wg *sync.WaitGroup, resultchan chan result) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed: %s\n", url)

	resultchan <- result{
		Value: url,
		Err:   nil,
	}
}

func Fanin2() {
	var wg sync.WaitGroup
	resultchan := make(chan result, 5)

	startTime := time.Now()
	wg.Add(5)
	go worker2("img_1.png", &wg, resultchan)
	go worker2("img_2.png", &wg, resultchan)
	go worker2("img_3.png", &wg, resultchan)
	go worker2("img_4.png", &wg, resultchan)
	go worker2("img_5.png", &wg, resultchan)

	wg.Wait()
	close(resultchan)

	for result := range resultchan {
		fmt.Printf("received: %s\n", result)
	}

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
