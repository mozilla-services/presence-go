package queue

import (
	"code.google.com/p/go-uuid/uuid"
	. "mozilla.org/presence"
)

// An Error represents a Storage backend error
type Error interface {
	error
	Unavailable() bool    // Is the storage system unavailable?
	RetryableError() bool // Was there an error likely temporary in nature?
	QueueFull() bool      // Is this queue full already?
}

// A QueueRetry indicates current queue related toggles for a message
type QueueRetry struct {
	CurrentTry int
	MaxTries   int // Maximum times to attempt delivery
	NextDelay  int // Used for various backoffs
}

type Queue interface {
	//  -- MCF Methods --

	// Returns a channel to send presence stanzas on for a given app
	PresenceDeliveryQueue(appId uuid.UUID) (chan<- *PresenceStanza, error)

	// -- Postmaster Methods --

	// Drain a queue of presence stanza's, retrieves N stanzas
	PresenceConsumerQueue(appId uuid.UUID) (<-chan *PresenceStanza, error)

	// Confirm handling of presence stanzas
	ConfirmPresenceStanzas(stanzaIds []uuid.UUID) error
}
