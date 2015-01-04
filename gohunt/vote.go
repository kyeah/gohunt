// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Vote struct {
	ID             int     `json:"id"`
	PostID         int     `json:"post_id"`
	Created        string  `json:"created_at"`
	User           User    `json:"user"`
}

func (v Vote) String() string {
	return fmt.Sprintf("vote[post %s: user %s]", v.PostID, v.User.Name)
}
