package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {

    accessToken := "ya29.GlxLBpEUgbAjSu_4LWRelYj2VpjUvAbXX2RlZWwdaMODPiL_rIPN7geyHnpBZZ3Z2Isj2XZLW043nPSLFmQrpCbpbjahvbk5_LCFsMwWrusPT8cjfHBVWz_izSz_UQ"

	// set paramater
    values := url.Values{}
    values.Add("timeMin", "2018-11-02T00:00:00-09:00")
    values.Add("timeMax", "2018-11-03T00:00:00-09:00")
	values.Add("fields",  "accessRole,defaultReminders,description,etag,items,kind,nextPageToken,nextSyncToken,summary,timeZone,updated")

	// create request url
	url := "https://www.googleapis.com/calendar/v3/calendars/gnavi.co.jp_2d38303539303533352d323836@resource.calendar.google.com/events?" + values.Encode()

	//
	req, _ := http.NewRequest("GET", url, nil)

	// set header
	req.Header.Set("Authorization", "OAuth " + accessToken)

	// call api
	client := new(http.Client)
	resp, err := client.Do(req)


    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()

    execute(resp)
}

func execute(response *http.Response) {
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}





//package main
//
//import (
//	"fmt"
//	"lib"
//)
//
//func main(){
//
//
//	// ユーザー設定情報取得
//	calendarIds, err := config.Parse("./config/calendar.json")
//	if err != nil {
//		fmt.Println("error ")
//	}
//
//	fmt.Println(calendarIds)
//
//}
