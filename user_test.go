package marksman

import (
	"context"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	root := os.Getenv("SNIPE_ROOT")
	if root == "" {
		t.Fatal("SNIPE_ROOT is not set")
	}

	token := os.Getenv("SNIPE_TOKEN")
	if token == "" {
		t.Fatal("SNIPE_TOKEN is not set")
	}

	client, err := New(
		WithRoot(root),
		WithToken(token),
	)
	if err != nil {
		t.Fatalf("Error creating client: %s", err.Error())
	}

	users, err := client.Me(context.Background())
	if err != nil {
		t.Fatalf("Error fetching users: %s", err.Error())
	}

	t.Log(users)
}
