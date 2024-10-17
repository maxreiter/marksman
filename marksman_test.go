package marksman

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/maxreiter/marksman/params/user"
	"github.com/maxreiter/marksman/snipeit"
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

		users, err := client.Users(ctx)
		assert.NilError(t, err, "fetching /users")

		t.Logf("Total users: %d", users.Total)
		for _, currentUser := range users.Rows {
			t.Logf("User: ID: %d: Name: %s", currentUser.ID, currentUser.Name)
		}

		meAgain, err := client.User(ctx, me.ID)
		assert.NilError(t, err, "fetching /users/%d", me.ID)

		t.Logf("User: ID: %d: Name: %s", meAgain.ID, meAgain.Name)

		err = client.CreateUser(
			ctx,
			user.FirstName("Dummy"),
			user.Username("dummy"),
			user.Password("insecure"),
		)
		assert.NilError(t, err, "creating new user")

		users, err = client.Users(ctx)
		assert.NilError(t, err, "fetching /users")

		var dummy *snipeit.User

		for _, currentUser := range users.Rows {
			if currentUser.FirstName == "Dummy" {
				dummy = currentUser
			}
		}

		err = client.UpdateUser(
			ctx,
			dummy.ID,
			user.FirstName("Dummy"),
			user.Username("dummy"),
			user.LastName("Test"),
		)
		assert.NilError(t, err, "updating user")

		dummy, err = client.User(ctx, dummy.ID)
		assert.NilError(t, err, "fetching dummy")

		assert.Assert(t, dummy.LastName == "Test", "dummy's last name is incorrect")

		err = client.DeleteUser(ctx, dummy.ID)
		assert.NilError(t, err, "deleting dummy")

		dummy, err = client.User(ctx, dummy.ID)
		t.Log(err)

	})
}
