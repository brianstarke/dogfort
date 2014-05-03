package domain

import (
	"testing"
	"time"
)

func TestAddAttachments(t *testing.T) {
	m := Message{"", "", "", "test text http://i.imgur.com/Or2tHrd.gif whatever", false, "", time.Now()}

	md := messageDomain{}

	md.addAttachments(&m)

	if !m.HasImage || len(m.Attachment) == 0 {
		t.Error("Message did not add image attachment")
	}
}
