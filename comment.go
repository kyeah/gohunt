// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Comment struct {
	ID               int
	PostID           int
	ParentCommentID  int
	UserID           int
	RepliesCount     int
	Body             string
	Created          float32
	Maker            bool	
	User             *User
	Replies          []*Comment
}

func (c *Comment) String() string {
	return fmt.Sprintf("%s: %s", c.User.Name, c.Body)
}
