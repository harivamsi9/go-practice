// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func returnType(url string, ch chan string) {

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("ERROR: ", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	ctype := resp.Header.Get("content-type")
// 	ch <- ctype
// 	// fmt.Println(ctype)
// }

// // func Serial(urls []string) {
// // 	now := time.Now()

// // 	for _, url := range urls {
// // 		returnType(url)
// // 	}

// // 	fmt.Println("Time Took: ", time.Since(now))
// // }

// // GO ROUTINES WAY
// func sitesConcurrent(urls []string) {
// 	now := time.Now()
// 	ch := make(chan string, len(urls))
// 	// var wg sync.WaitGroup

// 	for _, url := range urls {
// 		go returnType(url, ch)
// 		// wg.Add(1)
// 		// go func(url string) {
// 		// 	returnType(url)
// 		// 	wg.Done()
// 		// }(url)
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		val := <-ch
// 		fmt.Println(val)

// 	}
// 	// for val := range ch {
// 	// 	fmt.Println(val)
// 	// }
// 	close(ch)
// 	// wg.Wait()
// 	// close(ch)
// 	fmt.Println("Time Took: ", time.Since(now))

// }

// func main() {
// 	urls := []string{
// 		"https://google.com",
// 		"https://go.dev/",
// 		"https://www.meta.com/",
// 		"https://api.github.com",
// 	}

// 	// Serial(urls)
// 	sitesConcurrent(urls)

// }
