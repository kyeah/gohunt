// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"encoding/json"
	"net/url"
	"strconv"
)

var (
	base = "https://api.producthunt.com"
	postUrl = base + "/v1/posts"
)

type PostResponse struct {
	Posts  []Post  `json:"posts"`
}

func (c *Client) GetPosts() ([]Post, error) {
	return c.submitPostRequest(nil)
}

func (c *Client) GetPreviousPosts(daysAgo int) ([]Post, error) {
	values := &url.Values{
		"days_ago": { strconv.Itoa(daysAgo) },
	}
	return c.submitPostRequest(values)
}

func (c *Client) GetPostsOnDay(day string) ([]Post, error) {
	values := &url.Values{
		"day": { day },
	}
	return c.submitPostRequest(values)
}

func (c *Client) submitPostRequest(values *url.Values) ([]Post, error) {
	req := &Request{
		url: postUrl,
		values: values,
	}

	response, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	postmap := &PostResponse{}
	err = json.NewDecoder(response).Decode(postmap)
	if err != nil {
		return nil, err
	}

	return postmap.Posts, nil
}
