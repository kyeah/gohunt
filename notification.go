// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Notification struct {
	ID         int
	Type       string  // 'comment' or 'post'
	ShortBody  string
	FullBody   string
	Seen       bool
	Reference  *interface{}
	FromUser   *User
	ToUser     *User
}

func (n *Notification) String() string {
	return fmt.Sprintf("Notif: %s", n.FullBody)
}
