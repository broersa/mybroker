package dal

import "github.com/broersa/mybroker/dal"

type (
	// Dal Interface for Business Layer
	Dal interface {
		BeginTransaction() error
		CommitTransaction() error
		RollbackTransaction() error
		AddApplication(application *Application) (int64, error)
		GetApplication(id int64) (*Application, error)
		GetApplicationOnName(appname string) (*Application, error)
		GetApplicationOnAppEUI(appeui string) (*Application, error)
		AddDevice(device *Device) (int64, error)
		GetDevice(id int64) (*Device, error)
		GetDeviceOnDevEUI(deveui string) (*Device, error)
		GetDeviceOnAppEUIDevEUI(appeui string, deveui string) (*Device, error)
		AddSession(session *dal.Session) (int64, error)
		GetSessionOnID(id int64) (*Session, error)
		GetSessionOnDeviceActive(device int64) (*Session, error)
		GetActiveSessionOnDevEUIDevNonce(deveui string, devnonce string) (*Session, error)
		GetFreeNwkAddr() (uint32, error)
	}
)
