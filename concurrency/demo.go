package concurrency

import (
	"fmt"
	"time"
)

func worker(url string) string {
	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed: %s\n", url)
	return url
}

func Demo() {

	startTime := time.Now()

	result1 := worker("img_1.png")
	result2 := worker("img_2.png")

	fmt.Println("Results", result1)
	fmt.Println("Results", result2)

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
