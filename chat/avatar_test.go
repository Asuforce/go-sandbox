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

func TestGravatarAvatar(t *testing.T) {
	const (
		email      = "MyEmailAddress@example.com"
		correctURL = "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346"
	)

	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"email": email}

	if url, err := gravatarAvatar.GetAvatarURL(client); err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should return ErrNoAvatarURL when the value dosen't exist")
	} else if url != correctURL {
		t.Errorf("GravatarAvatar.GetAvatarURL returned the inccorect url of %s", url)
	}
}
