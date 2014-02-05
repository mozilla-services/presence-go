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
	// MCF Methods
	VerifyAppUidList(fxid string, appuids AppUidList) (valid bool, err error)
	LinkUuids(hostname net.IP, uids UidList) error
	// Unlink Uuids from this host if the user drops off
	// If the user was disconnected suddenly, indicate their zombie status
	UnlinkUuids(hostname net.IP, uids UidList, zombie bool) error
	GetLiveNotifications(uids UidList) ([]LiveNotification, error)
}
