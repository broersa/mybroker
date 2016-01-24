package dal

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
		//		GetSessionOnID(id int64) (*Session, error)
		//		GetSessionOnDevNonce(devnonce string) (*Session, error)
	}
)
