package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"./lib"
)

func main() {

	// get calendarid
	configs, _ := config.Parse("./config/calendar.json")

	// call api
	receiver := getInfo(configs.CalendarId)

	// parse
	for {
		receive, ok := <-receiver
		if !ok {
			log.Println("closed")
			return
		}
		log.Println(receive)
	}

}

func getInfo(urls map[string]string) <-chan string {

	receiver := make(chan string)

	wait := new(sync.WaitGroup)

	go func() {
		for mtgNo, url := range urls {

			// increment
			wait.Add(1)

			// call
			go func(url string) {

				res, err := http.Get(url)
				if err != nil {
					log.Fatal(err)
				}
				defer res.Body.Close()

				//
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				//
				receiver <- string(body)

				log.Println(mtgNo, res.Status)

				// decrement
				wait.Done()
			}(url)
		}
		wait.Wait()
		close(receiver)
	}()

	return receiver
}
