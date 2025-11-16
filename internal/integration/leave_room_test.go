package integration

import (
	"github.com/adnanahmady/go-websocket-chat/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeaveRoom(t *testing.T) {
	srv, err := test.Setup()
	assert.NoError(t, err)
	user1, err := srv.Connect("user1")
	assert.NoError(t, err)
	user2, err := srv.Connect("user2")
	assert.NoError(t, err)
	user3, err := srv.Connect("user3")
	assert.NoError(t, err)

	name := "given user3 when its connection closes then user1 and user2 should get leave message"
	t.Run(name, func(t *testing.T) {
		assert.NoError(t, user3.Close())

		assert.NoError(t, user1.SetReadDeadline(20))
		u1Msg := user1.ShouldReadType(t, "leave")
		assert.NoError(t, user2.SetReadDeadline(20))
		u2Msg := user2.ShouldReadType(t, "leave")

		assert.Equal(t, "user3", u1Msg["sender"].(map[string]any)["name"].(string))
		assert.Equal(t, "user3", u2Msg["sender"].(map[string]any)["name"].(string))
	})
}
