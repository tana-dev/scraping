package config

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

// 設定ファイルの値を表現する構造体
type Config struct {
	CalendarId string
}

// 設定ファイルを読み込む
func Parse(filename string) (Config, error) {

	var c Config

	fmt.Println("tanaka111111111")
	fmt.Println(filename)
	jsonString, err := ioutil.ReadFile(filename)
	fmt.Println(jsonString)
	if err != nil {
		fmt.Println("error: t1")
		return c, err
	}

	err = json.Unmarshal(jsonString, &c)
	if err != nil {
		fmt.Println("error: t1")
		return c, err
	}

	return c, nil
}
