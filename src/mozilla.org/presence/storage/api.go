package storage

import (
	"code.google.com/p/go-uuid/uuid"
	"net"
	"time"
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

type Storage interface {
	//  -- MCF Methods --

	// Verify a list of uid/appid's for a given FxID
	// The mapping is stored to ensure all UID's for an AppID can be
	// revoked if necessary.
	VerifyUidList(fxid string, appUids AppUidList) (valid bool, err error)

	// Link the list of uids/aids to the MCF's hostname
	LinkUids(hostname net.IP, uids UidList) error

	// Store a new Uid/Appid for a given FxID
	StoreUidForUser(fxid string, appUid AppUid) error

	// Unlink Uuids from this host if the user drops off
	// If the user was disconnected suddenly, zombie flag indicates the
	// uid should be added to the zombie queue *and* unlinked.
	// Version must match the version returned by HostnameForUid.
	// (Also used by Zombie Killer to evict a dead Uid)
	UnlinkUid(hostname net.IP, uid UserId, version int, zombie bool) error

	// Retrieve missed LiveNotifications for a batch of uids
	GetLiveNotifications(uids UidList) (notifs []LiveNotification, err error)

	// Delete live notifications
	DeleteLiveNotifications(uid UserId, messageIds []MessageId)

	// -- Postmaster Methods --

	// Retrieve the hostname for the UID
	// (Also used by Zombie Killer to verify Uid still has no hostname)
	HostnameForUid(uid uuid.UUID) (hostname net.IP, version int, err error)

	// Store a message for a Uid
	StoreLiveNotification(uid uuid.UUID, notif LiveNotification) (err error)

	// -- Zombie Killer Methods --

	// Retrieve a batch of Uid's from the zombie queue
	GetDeadUids(amount int) (uids []uuid.UUID, err error)

	// Remove Uid from Zombie queue
	RemoveDeadUid(uid uuid.UUID)

	// Get a list of userId's that have LiveNotifications waiting
	UsersWithLiveNotifications() (uids UidList, err error)

	// Get expired LiveNotifications
	GetOldLiveNotifications(uid UserId) (notifs []LiveNotification, err error)
}
