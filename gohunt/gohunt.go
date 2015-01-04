// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"encoding/json"
	"net/url"
)

var (
	base        = "https://api.producthunt.com"
	postUrl     = base + "/v1/posts"
	postAllUrl  = postUrl + "/all"
	postShowUrl = postUrl + "/show"
)

type SinglePostResponse struct {
	Post Post `json:"post"`
}

type MultiPostResponse struct {
	Posts []Post `json:"posts"`
}

func (c *Client) GetPost(id int) (Post, error) {
	values := &url.Values{
		"id": { string(id) },
	}
	return c.submitShowPostRequest(postShowUrl, values)
}

func (c *Client) GetPosts() ([]Post, error) {
	return c.submitPostRequest(postUrl, nil)
}

func (c *Client) GetPreviousPosts(daysAgo int) ([]Post, error) {
	values := &url.Values{
		"days_ago": { string(daysAgo) },
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
	if olderThanID > -1 { values.Add("older", string(olderThanID)) }
	if newerThanID > -1 { values.Add("newer", string(newerThanID)) }
	if count > -1       { values.Add("per_page", string(count))    }

	return c.submitPostRequest(postAllUrl, values)
}

func (c *Client) submitPostRequest(url string, values *url.Values) ([]Post, error) {
	postmap := &MultiPostResponse{}
	err := c.submitJsonRequest(url, values, postmap)
	if err != nil {
		return nil, err
	}
	return postmap.Posts, nil
}

func (c *Client) submitShowPostRequest(url string, values *url.Values) (Post, error) {
	postmap := &SinglePostResponse{}
	err := c.submitJsonRequest(url, values, postmap)
	if err != nil {
		return Post{}, err
	}
	return postmap.Post, nil
}

// Get a JSON Response using an arbitrary JSON template
func (c *Client) submitJsonRequest(url string, values *url.Values, jsonStruct interface{}) error {
	req := &Request{
		url: url,
		values: values,
	}
	
	response, err := c.sendRequest(req)
	if err != nil {
		return err
	}
	
	return json.NewDecoder(response).Decode(jsonStruct)
}
