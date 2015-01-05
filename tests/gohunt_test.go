package gohunt

import (
	"testing"
	"../gohunt"
)

var client *gohunt.Client

func initClient(t *testing.T) {
	if client == nil {
		var err error
		client = gohunt.NewUserClient(
			"devToken",
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

// Post Routes
func TestGetPost(t *testing.T) {
	initClient(t)
	_, err := client.GetPost(20)
	checkErr(t, err)
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


// User Routes
func TestGetUser(t *testing.T) {
	initClient(t)
	_, err := client.GetUser("kyeahokay")
	checkErr(t, err)
}

func TestGetAllUsers(t *testing.T) {
	initClient(t)
	users, err := client.GetAllUsers(500, 1000, 50, "asc")
	checkErr(t, err)
	checkArray(t, len(users))
}

// Vote Routes
func TestGetPostVotes(t *testing.T) {
	initClient(t)
	posts, err := client.GetPostVotes(10, 500, 1000, 100, "asc")
	checkErr(t, err)
	checkArray(t, len(posts))
}

func TestGetUserVotes(t *testing.T) {
	initClient(t)
	users, err := client.GetUserVotes(100, 500, 1000, 100, "asc")
	checkErr(t, err)
	checkArray(t, len(users))
}


// Notification Routes
func TestGetNotifications(t *testing.T) {
	initClient(t)
	_, err := client.GetNotifications(500, 1000, -1, "desc")
	checkErr(t, err)
}
