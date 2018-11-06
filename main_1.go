package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	receiver := worker()

	for {
		receive, ok := <-receiver
		if !ok {
			log.Println("closed")
			return
		}
		log.Println(receive)
	}
}

func worker() <-chan string {

	var wg sync.WaitGroup

	receiver := make(chan string)

	go func() {

		CalendarId := [...]string{"https://www.yahoo.co.jp/", "https://github.co.jp/"}

		for _, v := range CalendarId {

			// increment
			wg.Add(1)

			// goroutine
			go func(v string) {

				//
				req, _ := http.NewRequest("GET", v, nil)

				// call api
				client := new(http.Client)

				// call api
				resp, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()

				//
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				//
				receiver <- string(body)

				// decriment
				wg.Done()
			}(v)
		}
		wg.Wait()
		close(receiver)
	}()

	return receiver
}
