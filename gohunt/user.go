// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type User struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Username    string             `json:"username"`
	Headline    string             `json:"headline"`
	Created     string             `json:"created_at"`
	Image       map[string]string  `json:"image"`
	ProfileUrl  string             `json:"profile_url"`
	WebsiteUrl  string             `json:"website_url"`
}

func (u User) Summary() string {
	return fmt.Sprintf("user[%s: %s]", u.Name, u.Headline)
}
