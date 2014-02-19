package storage

import (
	"code.google.com/p/go-uuid/uuid"
	. "mozilla.org/presence"
	"net"
)

// An Error represents a Storage backend error
type Error interface {
	error
	Unavailable() bool    // Is the storage system unavailable?
	RetryableError() bool // Was there an error likely temporary in nature?
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

	// -- Postmaster Methods --

	// Retrieve the hostname for the UID
	// (Also used by Zombie Killer to verify Uid still has no hostname)
	HostnameForUid(uid uuid.UUID) (hostname net.IP, version int, err error)
}
