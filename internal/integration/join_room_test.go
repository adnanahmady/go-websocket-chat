package integration

import (
	"github.com/adnanahmady/go-websocket-chat/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinRoom(t *testing.T) {
	srv, err := test.Setup()
	assert.NoError(t, err)
	user1, err := srv.Connect("user1")
	assert.NoError(t, err)
	user2, err := srv.Connect("user2")
	assert.NoError(t, err)
	_ = user1.ShouldReadType(t, "join")

	name := "given user3 when its connected then join message should send to user1 and user2"
	t.Run(name, func(t *testing.T) {
		user3, err := srv.Connect("user3")
		assert.NoError(t, err)

		assert.NoError(t, user1.SetReadDeadline(20))
		u1Msg := user1.ShouldReadType(t, "join")
		assert.NoError(t, user2.SetReadDeadline(20))
		u2Msg := user2.ShouldReadType(t, "join")
		assert.NoError(t, user3.SetReadDeadline(20))
		_, err = user3.ReadType("join")
		assert.Errorf(t, err, "failed to assert user3 itself doesnt get join message")

		assert.Equal(t, "user3", u1Msg["sender"].(map[string]any)["name"].(string))
		assert.Equal(t, "user3", u2Msg["sender"].(map[string]any)["name"].(string))
	})
}
