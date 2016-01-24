package bll

type (
	// Bll Interface for Business Layer
	Bll interface {
		RegisterApplication(application *Application)
		GetApplicationOnAppEUI(appeui string) (*Application, error)
		ProcessJoinRequest(appeui string, deveui string, devnonce string)
	}
)
