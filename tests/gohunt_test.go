package gohunt

import (
	"testing"
	"../gohunt"
)

var client *gohunt.Client

func initClient(t *testing.T) {
	if client == nil {
		var err error
		client, err = gohunt.NewOAuthClient(
			"clientId",
			"clientSecret",
		)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("Generated Client")
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestGetPosts(t *testing.T) {
	initClient(t)
	_, err := client.GetPosts()
	checkErr(t, err)
}
