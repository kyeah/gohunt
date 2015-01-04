// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"encoding/json"
)

var (
	base = "https://api.producthunt.com"
	postUrl = base + "/v1/posts"
)

type PostResponse struct {
	Posts  []Post  `json:"posts"`
}

func (c *Client) GetPosts() ([]Post, error) {
	req := &Request{
		url: postUrl,
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
