package storage

import (
	"code.google.com/p/go-uuid/uuid"
	. "launchpad.net/gocheck"
	. "mozilla.org/presence"
	"net"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

type UserAppList map[string]AppUidList

type MemoryStorage struct {
	userApps UserAppList
}

func NewMemoryStorage() *MemoryStorage {
	apps := make(map[string]AppUidList)
	return &MemoryStorage{apps}
}

func sameAppUid(app1, app2 AppUid) bool {
	return uuid.Equal(app1.AppId, app2.AppId) && uuid.Equal(app1.Uid, app2.Uid)
}

func contains(applist AppUidList, element AppUid) bool {
	for _, existing := range applist {
		if sameAppUid(element, existing) {
			return true
		}
	}
	return false
}

func (s *MemoryStorage) VerifyUidList(fxid string, appuids AppUidList) (valid bool, err error) {
	for _, element := range appuids {
		if !contains(s.userApps[fxid], element) {
			return false, nil
		}
	}
	return true, nil
}

func (s *MemoryStorage) LinkUids(hostname net.IP, uids UidList) error {
	return nil
}

func (s *MemoryStorage) StoreUidForUser(fxid string, appUid AppUid) (err error) {
	s.userApps[fxid] = append(s.userApps[fxid], appUid)
	return nil
}

func (s *MemoryStorage) UnlinkUids(hostname net.IP, uids UidList, zombie bool) error {
	return nil
}

func (s *MemoryStorage) UnlinkUid(hostname net.IP, uid UserId, version int, zombie bool) error {
	return nil
}

func (s *MemoryStorage) GetLiveNotifications(uids UidList) ([]LiveNotification, error) {
	return nil, nil
}

func (s *MemoryStorage) HostnameForUid(uid uuid.UUID) (hostname net.IP, version int, err error) {
	return nil, -1, nil
}

func (s *MemoryStorage) StoreLiveNotification(uid uuid.UUID, notif LiveNotification) (err error) {
	return nil
}

func (s *MemoryStorage) GetDeadUids(amount int) (uids []uuid.UUID, err error) {
	return
}

func (s *MemoryStorage) RemoveDeadUid(uid uuid.UUID) {
	return
}

func (s *MemoryStorage) DeleteLiveNotifications(uid UserId, messageIds []MessageId) {
}

func (s *MemoryStorage) GetOldLiveNotifications(uid UserId) (notifs []LiveNotification, err error) {
	return
}

func (s *MemoryStorage) UsersWithLiveNotifications() (uids UidList, err error) {
	return
}

func (s *MySuite) TestStorageInterface(c *C) {
	c.Check(42, Equals, 42)
	memStorage := NewMemoryStorage()
	storage := Storage(memStorage)

	fxid := "tarek"
	uid, aid := uuid.NewUUID(), uuid.NewUUID()
	appuid := AppUid{uid, aid}

	// storing and uid for a given user id and app id
	storage.StoreUidForUser(fxid, appuid)

	// verify that we have that uid stored in memory
	appuids := []AppUid{AppUid{uid, aid}}
	result, _ := storage.VerifyUidList(fxid, appuids)

	c.Check(result, Equals, true)
}
