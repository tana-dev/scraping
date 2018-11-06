package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
//	"time"
	"lib"
	"sync"
	"testing"
)


func main() {

	// check access token

	// get calendarid
    //calendarId := [...]string{"gnavi.co.jp_31383538353532302d313232@resource.calendar.google.com", "gnavi.co.jp_2d3434303134363633333734@resource.calendar.google.com"}
	calendarIds, _ := config.Parse("./config/calendar.json")

	//
    result := testing.Benchmark(func(b *testing.B) { run(calendarIds) })
    fmt.Println(result.T)
}

func run(configs config.Config) {

	//
    fmt.Println("Start!")
    fmt.Println(len(configs.CalendarId))

    // WaitGroupを作成する
    wg := new(sync.WaitGroup)

    // channelを処理の数分だけ作成する
//    isFin := make(chan bool, len(configs.CalendarId))
//	m := new(sync.Mutex)
    isFin := make(chan bool, 1)

	// downloadセット
    for _, v := range configs.CalendarId {

//		m.Lock()
//		defer m.Unlock()

        // 処理1つに対して、1つ数を増やす
        wg.Add(1)

        // サブスレッドに処理を任せる
        go process(v, isFin, wg)
//        process(v)
    }

    // wg.Done
    wg.Wait()
    close(isFin)

	//
    fmt.Println("Finish!")

}

func process(id string, isFin chan bool, wg *sync.WaitGroup) {
//func process(id string) {

    // wgの数を1つ減らす
    defer wg.Done()

	// def
    accessToken := "ya29.GlxLBgJ3MDnFR7E2mbl2WZ9O6a5lB7BaAhoraC4wkp64AlMjXTkH4wG8KD4lY7GaIFH98h_WaVuJvS7dbwEco9C1QcDKfLdqsBTmW2ihoLMeJo2Kh_-SYTHDOtvn7Q"

	//
    fmt.Println(id)

	// set paramater
    values := url.Values{}
    values.Add("timeMin", "2018-11-02T00:00:00-09:00")
    values.Add("timeMax", "2018-11-03T00:00:00-09:00")
	values.Add("fields",  "accessRole,defaultReminders,description,etag,items,kind,nextPageToken,nextSyncToken,summary,timeZone,updated")

	// create request url
	url := "https://www.googleapis.com/calendar/v3/calendars/" + id + "/events?" + values.Encode()

	//
	req, _ := http.NewRequest("GET", url, nil)

	// set header
	req.Header.Set("Authorization", "OAuth " + accessToken)

	// call api
	client := new(http.Client)
	resp, err := client.Do(req)

//	time.Sleep(1 * time.Second)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

	//
    execute(resp)

	//
    isFin <- true

}

func execute(response *http.Response) {

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))

}






//	var i string
//
//	for i := 0; i < goroutines; i++{
//
//		// goroutines
//		go func(s chan<- int){
//
//			// def
//		    accessToken := "ya29.GlxJBiq5tL3WAJgefJjKeIZomlRz2HQzRyYvb6NWdVOnJVR1BW6AA6r8I6fulhIbUcw6vF8nywEgqThDHW-zzKA_NWnCYjqJ8VzVuzgJlQbsL9PZa2_q3LkgrjP3Lw"
//		    calendarId := "gnavi.co.jp_2d39393234313332342d353936@resource.calendar.google.com"
//
//			// set paramater
//		    values := url.Values{}
//		    values.Add("timeMin", "2018-11-02T00:00:00-09:00")
//		    values.Add("timeMax", "2018-11-03T00:00:00-09:00")
//			values.Add("fields",  "accessRole,defaultReminders,description,etag,items,kind,nextPageToken,nextSyncToken,summary,timeZone,updated")
//
//			// create request url
//			url := "https://www.googleapis.com/calendar/v3/calendars/" + calendarId + "/events?" + values.Encode()
//
//			req, _ := http.NewRequest("GET", url, nil)
//
//			// set header
//			req.Header.Set("Authorization", "OAuth " + accessToken)
//
//			// call api
//			client := new(http.Client)
//			resp, err := client.Do(req)
//		    if err != nil {
//		        fmt.Println(err)
//		        return
//		    }
//		    defer resp.Body.Close()
//
//		    execute(resp)
//			fmt.Println("処理完了")
//			s <- 0
//		}(c)
//
//	}
//
//	// wait
//	for i := 0; i < goroutines ; i++{
//		<-c
//	}
//
//	fmt.Println("すべて完了")
//
//
//
//
//}
//
//func getCalendarids(cids *[]string){
//}
//
//func execute(response *http.Response) {
//
//    body, err := ioutil.ReadAll(response.Body)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//    fmt.Println(string(body))
//
//}
