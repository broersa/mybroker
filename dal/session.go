package dal

import (
	"time"
)

// Session Data Entity
type Session struct {
	ID       int64
	Device   int64
	NwkAddr  int64
	DevNonce string
	AppNonce string
	NwkSKey  string
	AppSKey  string
	Active   time.Time
}
