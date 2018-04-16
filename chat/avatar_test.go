package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}

	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}

	if _, err := authAvatar.GetAvatarURL(testChatUser); err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when the value dosen't exist")
	}

	const testURL = "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser

	testUser.On("AvatarURL").Return(testURL, nil)

	if url, err := authAvatar.GetAvatarURL(testChatUser); err != nil {
		t.Error("AuthAvatar.GetAvatarURL should not return Error")
	} else if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	const (
		uniqueid   = "abc"
		correctURL = "//www.gravatar.com/avatar/abc"
	)

	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: uniqueid}

	if url, err := gravatarAvatar.GetAvatarURL(user); err != nil {
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
	user := &chatUser{uniqueID: userid}

	if url, err := fileSystemAvatar.GetAvatarURL(user); err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should return ErrNoAvatarURL when the value dosen't exist")
	} else if url != correctURL {
		t.Errorf("FileSystemAvatar.GetAvatarURL returned the inccorect url of %s", url)
	}
}
