package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"io/ioutil"
)


func sendRequest(url string, ch chan <- string)  {

	start := time.Now()
	response, _ := http.Get(url)

	seconds := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(response.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", seconds, len(body), url)
}

func main()  {

	start := time.Now()
 	ch := make(chan string)

 	for _, url := range os.Args[1:]{
 		go sendRequest(url, ch)
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
