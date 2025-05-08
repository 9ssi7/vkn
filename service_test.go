package vkn

import (
	"context"
	"testing"
)

func TestLogin(t *testing.T) {
	srv := New(Config{
		Username: "<GIB_USER_CODE>",
		Password: "<GIB_PW>",
	})
	if err := srv.Logout(context.Background()); err != nil {
		t.Fatal(err)
	}
	err := srv.Login(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
