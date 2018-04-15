package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	if _, err := authAvatar.GetAvatarURL(client); err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when the value dosen't exist")
	}

	const testURL = "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testURL}
	if url, err := authAvatar.GetAvatarURL(client); err != nil {
		t.Error("AuthAvatar.GetAvatarURL should not return Error")
	} else if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}
