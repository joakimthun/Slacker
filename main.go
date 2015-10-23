package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Token      string
	UserName   string
	Channel    string
	UserToPing string
}

func main() {
	config := getConfig("config.json")
	postMessage(config)
}

func postMessage(config Config) {
	url := fmt.Sprintf(
		"https://slack.com/api/chat.postMessage?token=%s&channel=%s&text=@%s&username=%s&link_names=1&pretty=1",
		config.Token, config.Channel, config.UserToPing, config.UserName)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func getConfig(fileName string) Config {
	buf, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(buf, &config)

	if err != nil {
		panic(err)
	}

	return config
}
