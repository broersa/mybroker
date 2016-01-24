package bll

type (
	// Bll Interface for Business Layer
	Bll interface {
		RegisterApplication(appname string) (int64, error)
		GetApplication(id int64) (*Application, error)
		GetApplicationOnAppEUI(appeui string) (*Application, error)
		ProcessJoinRequest(appeui string, deveui string, devnonce string)
	}
)
