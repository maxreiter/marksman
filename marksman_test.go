package marksman

import (
	"context"
	"os"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestMarksman(t *testing.T) {
	root := os.Getenv("MARKSMAN_TEST_ROOT")
	token := os.Getenv("MARKSMAN_TEST_TOKEN")

	assert.Assert(t, root != "", "$MARKSMAN_TEST_ROOT is not set")
	assert.Assert(t, token != "", "$MARKSMAN_TEST_TOKEN is not set")

	client, err := New(
		Root(root),
		Token(token),
	)
	assert.NilError(t, err, "creating client")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	t.Run("Test User Methods", func(t *testing.T) {
		me, err := client.Me(ctx)
		assert.NilError(t, err, "fetching /users/me")

		t.Logf("User: ID: %d: Name: %s", me.ID, me.Name)
	})
}
