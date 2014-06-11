package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func describe(o interface{}, tab int) {
	var buffer bytes.Buffer
	for i := 0; i < tab; i++ {
		buffer.WriteString("  ")
	}
	tabs := buffer.String()
	switch o.(type) {
	default:
		fmt.Printf("%vunexpected type %T", tabs, o) // %T prints whatever type t has
	case string:
		fmt.Printf("%v%v\n", tabs, o)
	case map[string]interface{}:
		m := o.(map[string]interface{})
		for k, _ := range m {
			fmt.Printf("%v%v\n", tabs, k)
			child := m[k]
			describe(child, tab+1)
		}
	case []interface{}:
		a := o.([]interface{})
		for i, child := range a {
			fmt.Printf("%v%v\n", tabs, i)
			describe(child, tab+1)
		}
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

	describe(info, 0)
}
