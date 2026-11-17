package connector

import (
	"testing"

	"maunium.net/go/mautrix/bridgev2"
	"maunium.net/go/mautrix/event"
)

func TestIsRelayedSticker(t *testing.T) {
	message := func(raw map[string]any, origSender *bridgev2.OrigSender) *bridgev2.MatrixMessage {
		return &bridgev2.MatrixMessage{MatrixEventBase: bridgev2.MatrixEventBase[*event.MessageEventContent]{
			Event:      &event.Event{Type: event.EventMessage, Content: event.Content{Raw: raw}},
			Content:    &event.MessageEventContent{MsgType: event.MsgImage},
			OrigSender: origSender,
		}}
	}
	relayed := &bridgev2.OrigSender{}
	sticker := map[string]any{"body": "Aguarde um momento", "url": "mxc://bessa.digital/x"}
	image := map[string]any{"msgtype": "m.image", "body": "foto.jpg", "url": "mxc://bessa.digital/y"}

	cases := []struct {
		name string
		msg  *bridgev2.MatrixMessage
		want bool
	}{
		{"relayed sticker", message(sticker, relayed), true},
		{"relayed image", message(image, relayed), false},
		{"sticker from the login owner", message(sticker, nil), false},
		{"content without raw", message(nil, relayed), false},
	}
	for _, c := range cases {
		if got := isRelayedSticker(c.msg); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}
