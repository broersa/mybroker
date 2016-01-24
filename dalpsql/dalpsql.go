package dalpsql

import (
	"database/sql"

	"github.com/broersa/mybroker/dal"
)

type (
	dalPsql struct {
		db *sql.DB
		tx *sql.Tx
	}
)

// New Implemented Factory
func New(db *sql.DB) dal.Dal {
	return &dalPsql{db, nil}
}

func (dalpsql *dalPsql) BeginTransaction() error {
	tx, err := dalpsql.db.Begin()
	if err != nil {
		return err
	}
	dalpsql.tx = tx
	return nil
}

func (dalpsql *dalPsql) CommitTransaction() error {
	err := dalpsql.tx.Commit()
	return err
}

func (dalpsql *dalPsql) RollbackTransaction() error {
	err := dalpsql.tx.Rollback()
	return err
}

func (dalpsql *dalPsql) AddApplication(application *dal.Application) (int64, error) {
	q, err1 := dalpsql.db.Prepare("insert into applications (appname, appeui) values ($1, $2) returning appkey")
	if err1 != nil {
		//		log.Fatal("here")
		return 0, err1
	}
	var r int64
	err2 := dalpsql.tx.Stmt(q).QueryRow(application.Name, application.AppEUI).Scan(&r)
	if err2 != nil {
		return 0, err2
	}
	return r, nil
}

func (dalpsql *dalPsql) GetApplicationOnAppEUI(appeui string) (*dal.Application, error) {
	var returnvalue dal.Application
	row := dalpsql.db.QueryRow("SELECT appkey, appname, appeui FROM applications WHERE appeui=$1", appeui)
	err := row.Scan(&returnvalue.ID, &returnvalue.Name, &returnvalue.AppEUI)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}
