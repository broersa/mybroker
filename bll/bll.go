package bll

type (
	// Bll Interface for Business Layer
	Bll interface {
		RegisterApplication(appname string) (int64, error)
		GetApplication(id int64) (*Application, error)
		GetApplicationOnAppEUI(appeui string) (*Application, error)
		RegisterDevice(appeui string, deveui string) (int64, error)
		GetDevice(id int64) (*Device, error)
		GetDeviceOnAppEUIDevEUI(appeui string, deveui string) (*Device, error)
		ProcessJoinRequest(appeui string, deveui string, devnonce string)
	}
)
