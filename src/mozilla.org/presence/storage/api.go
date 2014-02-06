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
type UidList []uuid.UUID

type LiveNotification struct {
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

	// Verify a list of AppUid's for a given FxID
	VerifyAppUidList(fxid string, appuids AppUidList) (valid bool, err error)

	// Link the list of uids to the MCF's hostname
	LinkUids(hostname net.IP, uids UidList) error

	// Store a Uid for an AppId for a given FxID
	StoreUidForUser(fxid string, uid, aid uuid.UUID) (err error)

	// Unlink Uuids from this host if the user drops off
	// If the user was disconnected suddenly, zombie flag indicates the
	// uid should be added to the zombie queue.
	// (Also used by Zombie Killer to evict a dead Uid)
	UnlinkUids(hostname net.IP, uids UidList, zombie bool) error

	// Retrieve missed LiveNotifications for a batch of uids
	GetLiveNotifications(uids UidList) ([]LiveNotification, error)

	// -- Postmaster Methods --

	// Retrieve the hostname for the UID
	// (Also used by Zombie Killer to verify Uid still has no hostname)
	HostnameForUid(uid uuid.UUID) (hostname net.IP, err error)

	// Store a message for a Uid
	StoreLiveNotification(uid uuid.UUID, notif LiveNotification) (err error)

	// -- Zombie Killer Methods --

	// Retrieve a batch of Uid's from the zombie queue
	GetDeadUids(amount int) (uids []uuid.UUID, err error)

	// Remove Uid from Zombie queue
	RemoveDeadUid(uid uuid.UUID)
}
