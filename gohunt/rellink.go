// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type RelatedLink struct {
	ID       int     `json:"id"`
	Url      string  `json:"url"`
	Title    string  `json:"title"`
	Domain   string  `json:"domain"`
	Favicon  string  `json:"favicon"`   
	PostID   int     `json:"post_id"`
	UserID   int     `json:"user_id"`
	Post     Post    `json:"post"`
}

func (c RelatedLink) Summary() string {
	return fmt.Sprintf("rellink[%s: %s]", c.Title, c.Url)
}
