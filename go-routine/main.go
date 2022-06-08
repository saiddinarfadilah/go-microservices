package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("initial goroutine :", runtime.NumGoroutine())
	defer fmt.Println("final goroutine :", runtime.NumGoroutine())

	start := time.Now()

	websites := map[string]string{
		"google":   "https://www.google.com",
		"youtube":  "https://www.youtube.com",
		"Bioskop":  "https://www.NontonMovie.com",
		"facebook": "https://www.facebook.com",
	}

	input := make(chan string)
	output := make(chan []string)

	var wg sync.WaitGroup

	defer close(output)

	go handleCheckWebsiteConcurrent(&wg, input, output)

	for i, v := range websites {
		wg.Add(1)
		go checkWebsite(i, v, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	for _, err := range result {
		if err != "" {
			fmt.Println(err)
		}
	}

	elapsed := time.Since(start)

	fmt.Println("Execution time : ", elapsed.String())

	return
}

func checkWebsite(t, w string, output chan string) {
	if _, err := http.Get(w); err != nil {
		output <- fmt.Sprintf("man, %s is down :(", t)
		return
	}

	fmt.Println(fmt.Sprintf("yeay, %s is up :)", t))
	output <- ""
}

func handleCheckWebsiteConcurrent(wg *sync.WaitGroup, input chan string, output chan []string) {
	err := make([]string, 0)

	for incomingErr := range input {
		err = append(err, incomingErr)
		wg.Done()
	}

	output <- err
}
