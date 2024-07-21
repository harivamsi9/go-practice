// /*
// Calculate the total download size for NYC taxi data for 2020

// For each month, we have two files: green and yellow. e.g:

// https://d37ci6vzurychx.cloudfront.net/trip-data/green_tripdata_2020-03.parquet
// https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2020-03.parquet

// Turn the below sequential algorithm to a concurrent one using goroutine per file!
// */

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// var (
// 	urlTemplate = "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2020-%02d.parquet"
// 	colors      = []string{"green", "yellow"}
// )

// func downloadSize(url string) (int, error) {
// 	// req, err := http.NewRequest(http.MethodHead, url, nil)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer resp.Body.Close()

// 	n := resp.Header.Get("content-length")
// 	size, err := strconv.Atoi(n)
// 	return size, nil

// }

// // func fastFetch(color string, month int, urlTemplate string, ch chan int) {
// // 	url := fmt.Sprintf(urlTemplate, color, month)
// // 	fmt.Println(url)
// // 	n, err := downloadSize(url)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	size, err := strconv.Atoi(n)
// // 	ch <- size //fmt.Println(size)
// // 	// return

// // }

// type result struct {
// 	url  string
// 	size int
// 	err  error
// }

// func sizeWorker(url string, ch chan result) {
// 	fmt.Println(url)
// 	res := result{url: url}
// 	res.size, res.err = downloadSize(url)
// 	ch <- res
// }

// func main() {
// 	start := time.Now()

// 	total_size := 0
// 	ch := make(chan result)

// 	for month := 1; month < 13; month++ {
// 		for _, color := range colors {

// 			// go fastFetch(color, month, url)
// 			// go fastFetch(color, month, urlTemplate, ch)

// 			url := fmt.Sprintf(urlTemplate, color, month)
// 			go sizeWorker(url, ch)

// 			// total_size += size
// 		}
// 	}

// 	for i := 0; i < 12*len(colors); i++ {

// 		size := <-ch
// 		total_size += size.size
// 	}

// 	fmt.Println(total_size)
// 	fmt.Println("Time Taken: ", time.Since(start))

// }

// // this is for (green, 03)
// // url := "https://d37ci6vzurychx.cloudfront.net/trip-data/green_tripdata_2020-03.parquet"
// // size, err := downloadSize(url)
// // if err != nil {
// // 	fmt.Println("ERROR: ", err)

// // }
// // 