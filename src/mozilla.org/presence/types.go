package presence

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

const (
	Online = iota
	Away
	Offline
)

type AppUid struct {
	AppId uuid.UUID
	Uid   uuid.UUID
}

type AppUidList []AppUid
type UserId uuid.UUID
type MessageId uuid.UUID
type UidList []UserId

type LiveNotification struct {
	// Unique ID for this LiveNotification
	MessageID uuid.UUID
	// AppID this message belongs to on the device
	AppId uuid.UUID
	// Passed to the application's LiveNotification handler
	Action []byte
	// Encrypted content payload the device must decrypt
	Message []byte
	// Latest point in time the message should be considered 'alive'
	TTL time.Time
}

type PresenceStanza struct {
	// Stanza ID
	StanzaId uuid.UUID
	// User ID and App ID for the stanza
	Origin UserId
	// Status type
	Status int
	// Custom message if any
	Message string
}
