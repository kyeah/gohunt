// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Post struct {
	ID            int               `json:"id,"`
	Name          string            `json:"name"`
	Tagline       string            `json:"tagline"`
	Created       string            `json:"created_at"`
	Day           string            `json:"day"`
	CommentsCount int               `json:"comments_count"`
	VotesCount    int               `json:"votes_count"`
	DiscussionUrl string            `json:"discussion_url"`
	RedirectUrl   string            `json:"redirect_url"`
	ScreenshotUrl map[string]string `json:"screenshot_url"`
	CurrentUser   currentUser       `json:"current_user"`
	User          User              `json:"user"`
	MakerInside   bool              `json:"maker_inside"`
	Makers        []User            `json:"makers"`
}

type currentUser struct {
	Voted     bool `json:"voted_for_post"`
	Commented bool `json:"commented_on_post"`
}

func (p Post) Summary() string {
	return fmt.Sprintf("post[%s: %s]", p.Name, p.Tagline)
}
