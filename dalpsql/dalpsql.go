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

func (dalpsql *dalPsql) GetApplication(id int64) (*dal.Application, error) {
	var returnvalue dal.Application
	row := dalpsql.db.QueryRow("SELECT appkey, appname, appeui FROM applications WHERE appkey=$1", id)
	err := row.Scan(&returnvalue.ID, &returnvalue.Name, &returnvalue.AppEUI)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) GetApplicationOnName(appname string) (*dal.Application, error) {
	var returnvalue dal.Application
	row := dalpsql.db.QueryRow("SELECT appkey, appname, appeui FROM applications WHERE appname=$1", appname)
	err := row.Scan(&returnvalue.ID, &returnvalue.Name, &returnvalue.AppEUI)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
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

func (dalpsql *dalPsql) AddDevice(device *dal.Device) (int64, error) {
	q, err1 := dalpsql.db.Prepare("insert into devices (devapp, deveui, devappkey) values ($1, $2, $3) returning devkey")
	if err1 != nil {
		//		log.Fatal("here")
		return 0, err1
	}
	var r int64
	err2 := dalpsql.tx.Stmt(q).QueryRow(device.Application, device.DevEUI, device.AppKey).Scan(&r)
	if err2 != nil {
		return 0, err2
	}
	return r, nil
}

func (dalpsql *dalPsql) GetDevice(id int64) (*dal.Device, error) {
	var returnvalue dal.Device
	row := dalpsql.db.QueryRow("SELECT devkey, devapp, deveui, devappkey FROM devices WHERE devkey=$1", id)
	err := row.Scan(&returnvalue.ID, &returnvalue.Application, &returnvalue.DevEUI, &returnvalue.AppKey)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) GetDeviceOnDevEUI(deveui string) (*dal.Device, error) {
	var returnvalue dal.Device
	row := dalpsql.db.QueryRow("SELECT devkey, devapp, deveui, devappkey FROM devices WHERE deveui=$1", deveui)
	err := row.Scan(&returnvalue.ID, &returnvalue.Application, &returnvalue.DevEUI, &returnvalue.AppKey)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}
