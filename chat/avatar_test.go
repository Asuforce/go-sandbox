package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

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
		userid     = "0bc83cb571cd1c50ba6f3e8a78ef1346"
		correctURL = "//www.gravatar.com/avatar/" + userid
	)

	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": userid}

	if url, err := gravatarAvatar.GetAvatarURL(client); err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should return ErrNoAvatarURL when the value dosen't exist")
	} else if url != correctURL {
		t.Errorf("GravatarAvatar.GetAvatarURL returned the inccorect url of %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	const (
		userid     = "abc"
		correctURL = "/avatars/abc.jpg"
	)

	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": userid}

	if url, err := fileSystemAvatar.GetAvatarURL(client); err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should return ErrNoAvatarURL when the value dosen't exist")
	} else if url != correctURL {
		t.Errorf("FileSystemAvatar.GetAvatarURL returned the inccorect url of %s", url)
	}
}
