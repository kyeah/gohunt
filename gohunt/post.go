// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Post struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Tagline        string  `json:"tagline"`
	Created        string  `json:"created_at"`
	Day            string  `json:"day"`
	CommentsCount  int     `json:"comments_count"`
	VotesCount     int     `json:"votes_count"`
	DiscussionUrl  string  `json:"discussion_url"`
	RedirectUrl    string  `json:"redirect_url"`
	ScreenshotUrl  string  `json:"screenshot_url"`
	Voted          bool    
	Commented      bool
	User           User    `json:"user"`
	MakerInside    bool    `json:"maker_inside"`
	Makers         []User  `json:"makers"`
}

func (p *Post) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.Tagline)
}
