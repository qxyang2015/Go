package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func mockSendToServer(url string) {
	fmt.Printf("server url: %s\n", url)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	urls := []string{"0.0.0.0:5000", "0.0.0.0:6000", "0.0.0.0:7000"}
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			mockSendToServer(url)
		}(url)
	}
	wg.Wait()
}
