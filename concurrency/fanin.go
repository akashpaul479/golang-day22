package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func worker1(url string, wg *sync.WaitGroup, resultchan chan string) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed: %s\n", url)

	resultchan <- url
}

func Fanin() {
	var wg sync.WaitGroup
	resultchan := make(chan string, 5)

	startTime := time.Now()
	wg.Add(5)
	go worker1("img_1.png", &wg, resultchan)
	go worker1("img_2.png", &wg, resultchan)
	go worker1("img_3.png", &wg, resultchan)
	go worker1("img_4.png", &wg, resultchan)
	go worker1("img_5.png", &wg, resultchan)

	wg.Wait()
	close(resultchan)

	for result := range resultchan {
		fmt.Printf("received: %s\n", result)
	}

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
