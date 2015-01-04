package gohunt

import (
	"testing"
	"../gohunt"
)

func genClient(t *testing.T) *gohunt.Client {
	client, err := gohunt.NewOAuthClient(
		"clientId",
		"clientSecret",
	)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestGetPosts(t *testing.T) {
	client := genClient(t)
	_, err := client.GetPosts()
	checkErr(t, err)
}
