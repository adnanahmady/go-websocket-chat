package integration

import (
	"github.com/adnanahmady/go-websocket-chat/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMessage(t *testing.T) {
	srv, err := test.Setup()
	assert.NoError(t, err)
	user1, err := srv.Connect("user1")
	assert.NoError(t, err)
	user2, err := srv.Connect("user2")
	assert.NoError(t, err)

	name := "given user1 when sends a message then user2 should receive it"
	t.Run(name, func(t *testing.T) {
		msg := test.Message{"type": "talk", "payload": "hello"}
		assert.NoError(t, user1.Send(msg))

		msg = user2.ShouldReadType(t, "talk")

		assert.Equal(t, "hello", msg["payload"])
	})

	name = "given message when has any type other than join or leave then receive talk type"
	t.Run(name, func(t *testing.T) {
		msg := test.Message{"type": "", "payload": "hello"}
		assert.NoError(t, user1.Send(msg))

		msg = user2.ShouldReadType(t, "talk")

		assert.Equal(t, "hello", msg["payload"])
	})
}
