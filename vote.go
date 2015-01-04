// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Vote struct {
	ID             int
	PostID         int
	Created        float32
	User           *User
}

func (p *Vote) String() string {
	return fmt.Sprintf("Vote on post %s: %s", v.PostID, v.User.Name)
}
