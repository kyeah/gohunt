// Copyright 2015 Kevin Yeh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gohunt

import (
	"fmt"
)

type PermissionInfo struct {
	CanVote     bool  `json:"can_vote_posts"`
	CanComment  bool  `json:"can_comment"`
	CanPost     bool  `json:"can_post"`
}

type NotificationInfo struct {
	Total   int  `json:"total"`
	Unseen  int  `json:"unseen"`
}

type UserSettings struct {
	ID                       int                `json:"id"`
	Name                     string             `json:"name"`
	Username                 string             `json:"username"`
	Headline                 string             `json:"headline"`
	Created                  string             `json:"created_at"`
	Image                    map[string]string  `json:"image"`
	ProfileUrl               string             `json:"profile_url"`
	WebsiteUrl               string             `json:"website_url"`
	VotesCount               int                `json:"votes_count"`
	PostsCount               int                `json:"posts_count"`
	MakerCount               int                `json:"maker_of_count"`
	FollowersCount           int                `json:"followers_count"`
	FollowingCount           int                `json:"followings_count"`
	SendMentionEmail         bool               `json:"send_mention_email"`
	SendMentionPush          bool               `json:"send_mention_push"`
	SendFriendPostEmail      bool               `json:"send_friend_post_email"`
	SendFriendPostPush       bool               `json:"send_friend_post_push"`
	PushSubscriber           bool               `json:"subscribed_to_push"`
	SendNewFollowerPush      bool               `json:"send_new_follower_push"`
	SendNewFollowerEmail     bool               `json:"send_new_follower_email"`
	SendAnnouncementPush     bool               `json:"send_announcement_push"`
	SendAnnouncementEmail    bool               `json:"send_announcement_email"`
	SendRecommendationPush   bool               `json:"send_recommendation_push"`
	SendRecommendationEmail  bool               `json:"send_recommendation_email"`
	Email                    string             `json:"email"`
	Role                     string             `json:"rile"`
	Permissions              PermissionInfo     `json:"permissions"`
	Notification             NotificationInfo   `json:"notifications"`
	FirstTimeUser            bool               `json:"first_time_user"`
	Votes                    []Vote             `json:"votes"`
	Posts                    []Post             `json:"posts"`
	MakerOf                  []Post             `json:"maker_of"`
	Followers                []User             `json:"followers"`
	Following                []User             `json:"followings"`
	
}

func (u UserSettings) Summary() string {
	return fmt.Sprintf("settings_for[%s: %s]", u.Name, u.Headline)
}
