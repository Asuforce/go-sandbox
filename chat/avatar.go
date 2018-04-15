package main

import "errors"

var ErrNoAvatarURL = errors.New("chat: Failed to get Avatar URL")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
