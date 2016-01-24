package bll

type (
	// Bll Interface for Business Layer
	Bll interface {
		RegisterApplication(application *Application)
		//		HasApplication(string appeui) (*Application, error)
		ProcessJoinRequest(appeui string, deveui string, devnonce string)
	}
)
