// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Post struct {
	ID             int
	Name           string
	Tagline        string
	Created        float32
	Day            string
	CommentsCount  int
	VotesCount     int
	DiscussionUrl  string
	RedirectUrl    string
	ScreenshotUrl  string
	Voted          bool
	Commented      bool
	User           *User
	MakerInside    bool
	Makers         []*User
}

func (p *Post) String() string {
	return fmt.Sprintf("%s: %s", p.Name, p.Tagline)
}
