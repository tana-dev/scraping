package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//	"time"
	"./lib"
	"log"
	"sync"
	"testing"
)

func main() {

	// check access token
	//checkAccessToken()

	// get calendarid
	calendarIds, _ := config.Parse("./config/calendar.json")

	// call api
	result := testing.Benchmark(func(b *testing.B) { run(calendarIds) })
	fmt.Println(result.T)

	// parse api result

	// echo
}

func run(configs config.Config) {

	// echo Start
	fmt.Println("Run Start!")

	// echo Number of Config Record
	fmt.Println(len(configs.CalendarId))

	// WaitGroupを作成する
	wg := new(sync.WaitGroup)

	// channelを処理の数分だけ作成する
	//	receiver:= make(chan bool, len(configs.CalendarId))
	m := new(sync.Mutex)
	receiver := make(chan string)

	//
	for _, v := range configs.CalendarId {

		m.Lock()
		defer m.Unlock()

		// 処理1つに対して、1つ数を増やす
		wg.Add(1)

		// サブスレッドに処理を任せる
		go process(v, receiver, wg)
	}

	// wg.Done
	wg.Wait()
	close(receiver)

	//
	log.Println(receiver)

	// echo Finish
	log.Println("Run Finish!")

}

func process(id string, receiver chan string, wg *sync.WaitGroup) {

	// wgの数を1つ減らす
	defer wg.Done()

	// def
	accessToken := "ya29"

	// set paramater
	values := url.Values{}
	values.Add("test", accessToken)
	//	values.Add("timeMin", "2018-11-02T00:00:00-09:00")
	//	values.Add("timeMax", "2018-11-03T00:00:00-09:00")
	//	values.Add("fields",  "accessRole,defaultReminders,description,etag,items,kind,nextPageToken,nextSyncToken,summary,timeZone,updated")

	// create request url
	//	url := "https://www.googleapis.com/calendar/v3/calendars/" + id + "/events?" + values.Encode()
	url := id + "?" + values.Encode()
	fmt.Println(url)

	//
	req, _ := http.NewRequest("GET", url, nil)

	// set header
	//	req.Header.Set("Authorization", "OAuth " + accessToken)

	// call api
	client := new(http.Client)
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
	//    fmt.Println(string(body))

	// parse api result
	//    parse(resp)

	receiver <- string(body)

}

func parse(response *http.Response) {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
