package main

import "testing"

func TestConfigure(t *testing.T) {
	c := config{APIKey: "your api key goes here", URL: "https://translate.yandex.net/api/v1.5/tr.json/translate?key="}
	p := c

	c.configure()

	if p != c {
		t.Error(`configure() method does not run correctly`)
	}
}
