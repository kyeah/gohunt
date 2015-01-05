// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Comment struct {
	ID               int        `json:"id"`
	PostID           int        `json:"post_id"`
	ParentCommentID  int        `json:"parent_comment_id"`
	UserID           int        `json:"user_id"`
	RepliesCount     int        `json:"child_comments_count"`
	Body             string     `json:"body"`
	Created          string     `json:"created_at"`
	Maker            bool	    `json:"maker"`
	User             User       `json:"user"`
	Replies          []Comment  `json:"child_comments"`
}

func (c Comment) Summary() string {
	return fmt.Sprintf("comment[%s: %s]", c.User.Name, c.Body)
}
