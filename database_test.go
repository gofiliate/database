package database

import (
	"testing"
)

func TestGetConnectionString(t *testing.T) {

	want := "username:password@tcp(host:3306)/database"
	got := GetConnectionString("username", "password", "host", 3306, "database")

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
