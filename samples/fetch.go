package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/timburks/agents"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	response, err := http.Get("https://itunes.apple.com/us/rss/topfreeapplications/limit=5/json")
	check(err)

	contents, err := ioutil.ReadAll(response.Body)
	check(err)

	var info map[string]interface{}
	err = json.Unmarshal(contents, &info)
	check(err)

	//var feed = info["feed"].(map[string]interface{})

	agents.Describe(info, 0)
}
