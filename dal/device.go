package dal

// Device Data Entity
type Device struct {
	ID          int64
	Application int64
	DevEUI      string
	AppKey      string
}
