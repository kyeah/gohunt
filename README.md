Gohunt
========

A golang client library for the official Product Hunt API.

## Usage

Download and install gohunt by running `go get github.com/kyeah/gohunt/gohunt`.

```go
package main

import (
	"fmt"
	"log"
	"github.com/kyeah/gohunt/gohunt"
)

func main() {
	client := gohunt.NewUserClient("devToken")

	// Grab today's posts
	posts, err := client.GetPosts()
	if err != nil {
		log.Fatal(err)
	}

        // Print post summaries (Title: headline)
	for _, post := range posts {
		fmt.Println(post.Summary())
	}
}
```

### Gohunt Client

Interaction with the Product Hunt API is facilitated by the Gohunt Client. The client can be generated in three ways:


Client-Only Authentication by OAuth2
```go
client, err := gohunt.NewOAuthClient(clientID, clientSecret)
```

User-Authentication by Developer Token
```go
client := gohunt.NewUserClient(phToken)
```

User-Authentication by OAuth2
```go
func HandleLogin() {
   err := gohunt.RequestUserOAuthCode(clientID, redirectUrl, state)
}

func HandleRedirect(data) {
   client, err := gohunt.NewUserOAuthClient(clientID, clientSecret, redirectUrl, data.code)
}
```

### Client Interface

`[param]` indicates an optional parameter; use "" or -1 to exclude them in the request.

```go
// Posts
client.GetPost(id int)
client.GetPosts()
client.GetPreviousPosts(daysAgo int)
client.GetPostsOnDay(day string) // Formatted YYYY-MM-DD
client.GetAllPosts([searchUrl string], [olderThanID int], [newerThanID int], [count int])

// Users
client.GetUser(username string)  // id or username
client.GetAllUsers([olderThanID int], [newerThanID int], [count int], [order string])  // order is "asc" or "desc"

// Votes
client.GetPostVotes(postID int, [olderThanID int], [newerThanID int], [count int], [order string])
client.GetUserVotes(userID int, [olderThanID int], [newerThanID int], [count int], [order string])

// Comments
client.GetPostComments(postID int, [olderThanID int], [newerThanID int], [count int], [order string])
client.GetUserComments(userID int, [olderThanID int], [newerThanID int], [count int], [order string])

// Followers and Following
client.GetFollowers(userID int, [olderThanID int], [newerThanID int], [count int], [order string])
client.GetFollowing(userID int, [olderThanID int], [newerThanID int], [count int], [order string])

// Related Links
client.GetRelatedLinks([searchUrl string])

// Requires User-Authenticated Client
client.GetNotifications([olderThanID int], [newerThanID int], [count int], [order string])
client.GetSettings()
```
