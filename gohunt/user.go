// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type User struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	Username         string            `json:"username"`
	Headline         string            `json:"headline"`
	Created          string            `json:"created_at"`
	ImageUrl         map[string]string `json:"image_url"`
	ProfileUrl       string            `json:"profile_url"`
	WebsiteUrl       string            `json:"website_url"`
	Votes            []Vote            `json:"votes"`
	Posts            []Post            `json:"posts"`
	MakerOf          []Post            `json:"maker_of"`
	Followers        []User            `json:"followers"`
	Following        []User            `json:"followings"`
	CollectionsCount int64             `json:"collections_count"`
	CreatedAt        string            `json:"created_at"`
	FollowersCount   int64             `json:"followers_count"`
	FollowingsCount  int64             `json:"followings_count"`
	MakerOfCount     int64             `json:"maker_of_count"`
	PostsCount       int64             `json:"posts_count"`
	VotesCount       int64             `json:"votes_count"`
}

func (u User) Summary() string {
	return fmt.Sprintf("user[%s: %s]", u.Name, u.Headline)
}
