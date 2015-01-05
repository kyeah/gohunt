// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type Notification struct {
	ID         int          `json:"id"`
	Type       string       `json:"type"`  // 'comment' or 'post'
	ShortBody  string       `json:"body"`
	FullBody   string       `json:"sentence"`
	Seen       bool         `json:"seen"`
	Reference  interface{}
	FromUser   User         `json:"from_user"`
	ToUser     User         `json:"to_user"`
}

func (n Notification) Summary() string {
	return fmt.Sprintf("notif[%s]", n.FullBody)
}
