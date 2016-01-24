package dal

// Session Data Entity
type Session struct {
	ID       int64
	Device   int64
	DevNonce string
	AppNonce string
}
