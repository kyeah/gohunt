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

func checkArray(t *testing.T, length int ) {
	if length == 0 {
		t.Log("No elements in array.")
		t.Fail()
	}
}

func TestGetPosts(t *testing.T) {
	initClient(t)
	posts, err := client.GetPosts()
	checkErr(t, err)
	checkArray(t, len(posts))
}

func TestGetPreviousPosts(t *testing.T) {
	initClient(t)
	posts, err := client.GetPreviousPosts(5)
	checkErr(t, err)
	checkArray(t, len(posts))
}

func TestGetPostsOnDay(t *testing.T) {
	initClient(t)
	posts, err := client.GetPostsOnDay("2014-12-25")
	checkErr(t, err)
	checkArray(t, len(posts))
}

func TestGetAllPosts(t *testing.T) {
	initClient(t)
	posts, err := client.GetAllPosts("", 500, 1000, 3)
	checkErr(t, err)
	checkArray(t, len(posts))
}
