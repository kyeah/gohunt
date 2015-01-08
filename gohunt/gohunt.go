// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

var (
	base           = "https://api.producthunt.com"
	postUrl        = base + "/v1/posts"
	userUrl        = base + "/v1/users"
	notifUrl       = base + "/v1/notifications"
	relLinkUrl     = base + "/v1/related_links"
	commentUrl     = base + "/v1/comments"
	settingsUrl    = base + "/v1/me"
	postAllUrl     = postUrl + "/all"
	postVoteUrl    = postUrl + "/%s/votes"
	userVoteUrl    = userUrl + "/%s/votes"
	postCommentUrl = postUrl + "/%s/comments"
	userCommentUrl = userUrl + "/%s/comments"
	followerUrl    = userUrl + "/%s/followers"
	followingUrl   = userUrl + "/%s/following"
)

type singlePostResponse struct {
	Post Post `json:"post"`
}

type multiPostResponse struct {
	Posts []Post `json:"posts"`
}

type singleUserResponse struct {
	User User `json:"user"`
}

type multiUserResponse struct {
	Users []User `json:"users"`
}

type notificationResponse struct {
	Notifs []Notification `json:"notifications"`
}

type singleVoteResponse struct {
	Vote Vote `json:"vote"`
}

type voteResponse struct {
	Votes []Vote `json:"votes"`
}

type singleCommentResponse struct {
	Comment Comment `json:"comment"`
}

type commentResponse struct {
	Comments []Comment `json:"comments"`
}

type followInnerData struct {
	ID   int `json:"id"`
	User User `json:"user"`
}

type followersResponse struct {
	Data []followInnerData `json:"followers"`
}

type followingResponse struct {
	Data []followInnerData `json:"following"`
}

type settingsResponse struct {
	Settings UserSettings `json:"user"`
}

type relatedLinkResponse struct {
	Links []RelatedLink `json:"related_links"`
}

type errorResponse struct {
	Error       string
	Description string
}

// Post Routes
func (c *Client) GetPost(id int) (Post, error) {
	return c.submitSinglePostRequest(postUrl + "/" + strconv.Itoa(id), nil)
}

func (c *Client) GetPosts() ([]Post, error) {
	return c.submitPostRequest(postUrl, nil)
}

func (c *Client) GetPreviousPosts(daysAgo int) ([]Post, error) {
	values := &url.Values{
		"days_ago": { strconv.Itoa(daysAgo) },
	}
	return c.submitPostRequest(postUrl, values)
}

func (c *Client) GetPostsOnDay(day string) ([]Post, error) {
	values := &url.Values{
		"day": { day },
	}
	return c.submitPostRequest(postUrl, values)
}

func (c *Client) GetAllPosts(searchUrl string, olderThanID int, newerThanID int, count int) ([]Post, error) {	
	values := &url.Values{}
	if searchUrl != ""  { values.Add("search[url]", searchUrl)     }
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	
	return c.submitPostRequest(postAllUrl, values)
}

func (c *Client) CreatePost(link string, name string, tagline string) (Post, error) {
	values := &url.Values{
		"action": { "POST" },
		"url": { link },
		"name": { name },
		"tagline": { tagline },
	}
	return c.submitSinglePostRequest(postUrl, values)
}

func (c *Client) submitPostRequest(url string, values *url.Values) ([]Post, error) {
	postmap := &multiPostResponse{}
	err := c.submitJsonRequest(url, values, postmap)
	if err != nil {
		return nil, err
	}
	return postmap.Posts, nil
}

func (c *Client) submitSinglePostRequest(url string, values *url.Values) (Post, error) {
	postmap := &singlePostResponse{}
	err := c.submitJsonRequest(url, nil, postmap)
	if err != nil {
		return Post{}, err
	}
	return postmap.Post, nil
}


// User Routes
func (c *Client) GetUser(username string) (User, error) {
	return c.submitShowUserRequest(userUrl + "/" + username)
}

func (c *Client) GetAllUsers(olderThanID int, newerThanID int, count int, order string) ([]User, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	return c.submitUserRequest(userUrl, values)
}

func (c *Client) submitUserRequest(url string, values *url.Values) ([]User, error) {
	usermap := &multiUserResponse{}
	err := c.submitJsonRequest(url, values, usermap)
	if err != nil {
		return nil, err
	}
	return usermap.Users, nil
}

func (c *Client) submitShowUserRequest(url string) (User, error) {
	usermap := &singleUserResponse{}
	err := c.submitJsonRequest(url, nil, usermap)
	if err != nil {
		return User{}, err
	}
	return usermap.User, nil
}


// Vote Routes
func (c *Client) GetPostVotes(postID int, olderThanID int, newerThanID int, count int, order string) ([]Vote, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	id := strconv.Itoa(postID)
	return c.submitVoteRequest(fmt.Sprintf(postVoteUrl, id), values)
}


func (c *Client) GetUserVotes(userID int, olderThanID int, newerThanID int, count int, order string) ([]Vote, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	id := strconv.Itoa(userID)
	return c.submitVoteRequest(fmt.Sprintf(userVoteUrl, id), values)
}

func (c *Client) VoteForPost(postID int, voting bool) (Vote, error) {
	var action string
	if voting {
		action = "POST"
	} else {
		action = "DELETE"
	}
	values := &url.Values{ 
		"action": { action },
		"post_id": { strconv.Itoa(postID) },
	}

	id := strconv.Itoa(postID)
	votemap := &singleVoteResponse{}
	err := c.submitJsonRequest(fmt.Sprintf(postVoteUrl, id), values, votemap)
	if err != nil {
		return Vote{}, err
	}
	return votemap.Vote, nil
}

func (c *Client) submitVoteRequest(url string, values *url.Values) ([]Vote, error) {
	votemap := &voteResponse{}
	err := c.submitJsonRequest(url, values, votemap)
	if err != nil {
		return nil, err
	}
	return votemap.Votes, nil
}


// Comment Routes
func (c *Client) GetPostComments(postID int, olderThanID int, newerThanID int, count int, order string) ([]Comment, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	id := strconv.Itoa(postID)
	return c.submitCommentRequest(fmt.Sprintf(postCommentUrl, id), values)
}


func (c *Client) GetUserComments(userID int, olderThanID int, newerThanID int, count int, order string) ([]Comment, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	id := strconv.Itoa(userID)
	return c.submitCommentRequest(fmt.Sprintf(userCommentUrl, id), values)
}

func (c *Client) CreateComment(postID int, parentCommentID int, body string) (Comment, error) {
	values := &url.Values{
		"action": { "POST" },
		"body": { body },
	}
	if parentCommentID != -1 { 
		values.Add("parent_comment_id", strconv.Itoa(parentCommentID))
	}
	id := strconv.Itoa(postID)
	return c.submitSingleCommentRequest(fmt.Sprint(postCommentUrl, id), values)
}

func (c *Client) UpdateComment(commentID int, parentCommentID int, body string) (Comment, error) {
	values := &url.Values{
		"action": { "PUT" },
		"body": { body },
	}
	if parentCommentID != -1 { 
		values.Add("parent_comment_id", strconv.Itoa(parentCommentID))
	}
	id := strconv.Itoa(commentID)
	return c.submitSingleCommentRequest(commentUrl + "/" + id, values)
}

func (c *Client) submitCommentRequest(url string, values *url.Values) ([]Comment, error) {
	commentmap := &commentResponse{}
	err := c.submitJsonRequest(url, values, commentmap)
	if err != nil {
		return nil, err
	}
	return commentmap.Comments, nil
}

func (c *Client) submitSingleCommentRequest(url string, values *url.Values) (Comment, error) {
	commentmap := &singleCommentResponse{}
	err := c.submitJsonRequest(url, values, commentmap)
	if err != nil {
		return Comment{}, err
	}
	return commentmap.Comment, nil
}


// Notification Routes
func (c *Client) GetNotifications(olderThanID int, newerThanID int, count int, order string) ([]Notification, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	return c.submitNotificationRequest(notifUrl, values)
}

func (c *Client) ClearNotifications() ([]Notification, error) {
	values := &url.Values {
		"action": { "DELETE" },
	}
	return c.submitNotificationRequest(notifUrl, values)
}

func (c *Client) submitNotificationRequest(url string, values *url.Values) ([]Notification, error) {
	notifmap := &notificationResponse{}
	err := c.submitJsonRequest(url, values, notifmap)
	if err != nil {
		return nil, err
	}
	return notifmap.Notifs, nil
}


// Follow Routes
func (c *Client) GetFollowers(userID int, olderThanID int, newerThanID int, count int, order string) ([]User, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	id := strconv.Itoa(userID)
	return c.submitFollowersRequest(fmt.Sprintf(followerUrl, id), values)
}

func (c *Client) GetFollowing(userID int, olderThanID int, newerThanID int, count int, order string) ([]User, error) {	
	values := &url.Values{}
	if olderThanID > -1 { values.Add("older", strconv.Itoa(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", strconv.Itoa(newerThanID)) }
	if count > -1       { values.Add("per_page", strconv.Itoa(count))    }
	if order != ""      { values.Add("order", order)                     }

	id := strconv.Itoa(userID)
	return c.submitFollowingRequest(fmt.Sprintf(followingUrl, id), values)
}

func (c *Client) submitFollowersRequest(url string, values *url.Values) ([]User, error) {
	usermap := &followersResponse{}
	err := c.submitJsonRequest(url, values, usermap)
	if err != nil {
		return nil, err
	}
	
	users := make([]User, len(usermap.Data))
	
	for i := 0; i < len(usermap.Data); i++ {
		users[i] = usermap.Data[i].User
	}
	return users, nil
}

func (c *Client) submitFollowingRequest(url string, values *url.Values) ([]User, error) {
	usermap := &followingResponse{}
	err := c.submitJsonRequest(url, values, usermap)
	if err != nil {
		return nil, err
	}
	
	users := make([]User, len(usermap.Data))
	
	for i := 0; i < len(usermap.Data); i++ {
		users[i] = usermap.Data[i].User
	}
	return users, nil
}


// Related Links Route
func (c *Client) GetRelatedLinks(searchUrl string) ([]RelatedLink, error) {
	values := &url.Values{}
	if searchUrl != ""  { values.Add("search[url]", searchUrl) }
	
	linkmap := &relatedLinkResponse{}
	err := c.submitJsonRequest(relLinkUrl, values, linkmap)
	if err != nil {
		return nil, err
	}
	return linkmap.Links, nil
}


// Settings Route
func (c *Client) GetSettings() (UserSettings, error) {
	setmap := &settingsResponse{}
	err := c.submitJsonRequest(settingsUrl, nil, setmap)
	if err != nil {
		return UserSettings{}, err
	}
	return setmap.Settings, nil
}


// Get a JSON Response using an arbitrary JSON template
func (c *Client) submitJsonRequest(url string, values *url.Values, jsonStruct interface{}) error {
	req := &Request{
		url: url,
		values: values,
	}
	
	response, err := c.sendRequest(req)
	if err != nil {
		errorstruct := &errorResponse{}
		err = json.NewDecoder(response).Decode(errorstruct)
		if err == nil {
			err = errors.New(fmt.Sprintf("[%s] %s", errorstruct.Error, errorstruct.Description))
		}
		return err
	}

	err = json.NewDecoder(response).Decode(jsonStruct)
	return err
}
