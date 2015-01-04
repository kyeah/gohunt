// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type User struct {
	ID          int
	Name        string
	Username    string
	Headline    string
	Created     float32
	Image       map[string]string
	ProfileUrl  string
	WebsiteUrl  string
}

func (u *User) String() string {
	return fmt.Sprintf("%s (%s)", u.Name, u.Headline)
}
