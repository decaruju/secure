package model

import (
	"testing"
)

func TestPasswordFlow(t *testing.T) {
	username := "Some username"
	password := "Some password"
	newPassword := "Some new password"
	wrongPassword := "Some wrong password"

	user := User{Username: username}
	err := user.CreatePassword(password)
	if err != nil {
		t.Errorf("Password create failed")
	}
	err = user.CheckPassword(wrongPassword)
	if err == nil {
		t.Errorf("Wrong password checked successfully")
	}
	err = user.CheckPassword(password)
	if err != nil {
		t.Errorf("Password check failed")
	}
	err = user.UpdatePassword(wrongPassword, newPassword)
	if err == nil {
		t.Errorf("Wrong password updated new password")
	}
	err = user.UpdatePassword(password, newPassword)
	if err != nil {
		t.Errorf("Password update failed")
	}
	err = user.CheckPassword(password)
	if err == nil {
		t.Errorf("Old password checked successfully, update failed")
	}
	err = user.CheckPassword(newPassword)
	if err != nil {
		t.Errorf("New password did not work, update failed")
	}
}
