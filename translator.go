package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type responseJSON struct {
	Code    int      `json:"code"`
	Lang    string   `json:"lang"`
	Text    []string `json:"text"`
	Message string   `json:"message,omitempty"`
}

type config struct {
	URL    string
	APIKey string
}

func (c *config) configure() error {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
		showErrorAndExit(fmt.Sprintf("file read error: %v", err), false)
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		showErrorAndExit(fmt.Sprintf("unmarshal error: %v", err), false)
	}
	return nil
}

var lang = flag.String("l", "en-tr", "translation direction, from-to (en-tr)")
var text = flag.String("t", "", "text to be translated")

var usage = `Usage: translator -t Hello -l en-tr 

	-t Text to be translated. If text contains space, quotation marks must be used
	-l Translation direction, default value is en-tr which means translator translates text from English to Turkish 
	
	`

func showErrorAndExit(message string, aboutUsage bool) {
	fmt.Fprintf(os.Stderr, message)
	if aboutUsage {
		flag.Usage()
	}
	os.Exit(1)
}

func getResponseAsJSON(c config) responseJSON {

	var data = make(url.Values)
	data["text"] = []string{*text}

	var url = fmt.Sprintf("%s%s&lang=%s", c.URL, c.APIKey, *lang)

	resp, err := http.PostForm(url, data)

	if err != nil {
		showErrorAndExit(fmt.Sprintf("Request error: %v", err), false)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		showErrorAndExit(fmt.Sprintf("Response error: %v\n", err), false)
	}

	var respJSON responseJSON
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		showErrorAndExit(fmt.Sprintf("Unmarshal error: %v\n", err), false)
	}

	return respJSON
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()
	if *text == "" {
		showErrorAndExit("Usage error", true)
	}

	c := config{}

	err := c.configure()
	if err != nil {
		showErrorAndExit(fmt.Sprintf("config error: %v", err), false)
	}

	response := getResponseAsJSON(c)
	if response.Code == 200 {
		fmt.Printf("%s: %s\n", *text, response.Text[:])
	} else {
		showErrorAndExit(fmt.Sprintf("Error: %s\n", response.Message), false)
	}

}
