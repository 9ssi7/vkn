package vkn

import (
	"context"
	"testing"
)

func TestLogin(t *testing.T) {
	srv := New(Config{
		Username: "46915028",
		Password: "549701",
	})
	if err := srv.Logout(context.Background()); err != nil {
		t.Fatal(err)
	}
	err := srv.Login(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
