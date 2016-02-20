package dalpsql

import (
	"database/sql"
	"fmt"

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
	fmt.Println(appeui)
	var returnvalue dal.Application
	row := dalpsql.db.QueryRow("SELECT appkey, appname, appeui FROM applications WHERE appeui=$1", appeui)
	err := row.Scan(&returnvalue.ID, &returnvalue.Name, &returnvalue.AppEUI)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	fmt.Println("andre")
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

func (dalpsql *dalPsql) GetDeviceOnAppEUIDevEUI(appeui string, deveui string) (*dal.Device, error) {
	var returnvalue dal.Device
	row := dalpsql.db.QueryRow("SELECT devkey, devapp, deveui, devappkey FROM devices JOIN applications on devapp=appkey WHERE appeui=$1 AND deveui=$2", appeui, deveui)
	err := row.Scan(&returnvalue.ID, &returnvalue.Application, &returnvalue.DevEUI, &returnvalue.AppKey)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) GetSessionOnID(id int64) (*dal.Session, error) {
	var returnvalue dal.Session
	row := dalpsql.db.QueryRow("SELECT seskey, sesdev, sesdevnonce, sesappnonce, sesnwkaddr, sesnwkskey, sesappskey, sesactive FROM sessions WHERE seskey=$1", id)
	err := row.Scan(&returnvalue.ID, &returnvalue.Device, &returnvalue.DevNonce, &returnvalue.AppNonce, &returnvalue.NwkAddr, &returnvalue.NwkSKey, &returnvalue.AppSKey, &returnvalue.Active)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) GetSessionOnDeviceActive(device int64) (*dal.Session, error) {
	var returnvalue dal.Session
	row := dalpsql.db.QueryRow("SELECT seskey, sesdev, sesdevnonce, sesappnonce, sesnwkaddr, sesnwkskey, sesappskey, sesactive FROM sessions WHERE sesdev=$1 and sesactive > now()", device)
	err := row.Scan(&returnvalue.ID, &returnvalue.Device, &returnvalue.DevNonce, &returnvalue.AppNonce, &returnvalue.NwkAddr, &returnvalue.NwkSKey, &returnvalue.AppSKey, &returnvalue.Active)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) GetSessionOnDeviceDevNonceActive(device int64, devnonce string) (*dal.Session, error) {
	var returnvalue dal.Session
	row := dalpsql.db.QueryRow("SELECT seskey, sesdev, sesdevnonce, sesappnonce, sesnwkaddr, sesnwkskey, sesappskey, sesactive FROM sessions WHERE sesdev=$1 and sesdevnonce=$2 and sesactive > now()", device, devnonce)
	err := row.Scan(&returnvalue.ID, &returnvalue.Device, &returnvalue.DevNonce, &returnvalue.AppNonce, &returnvalue.NwkAddr, &returnvalue.NwkSKey, &returnvalue.AppSKey, &returnvalue.Active)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) AddSession(session *dal.Session) (int64, error) {
	q, err1 := dalpsql.db.Prepare("insert into sessions (sesdev, sesdevnonce, sesappnonce, sesnwkaddr, sesnwkskey, sesappskey, sesactive) values ($1, $2, $3, $4, $5, $6, $7) returning seskey")
	if err1 != nil {
		return 0, err1
	}
	var r int64
	err2 := dalpsql.tx.Stmt(q).QueryRow(session.Device, session.DevNonce, session.AppNonce, session.NwkAddr, session.NwkSKey, session.AppSKey, session.Active).Scan(&r)
	if err2 != nil {
		return 0, err2
	}
	return r, nil
}

func (dalpsql *dalPsql) GetFreeNwkAddr() (*dal.NwkAddr, error) {
	var returnvalue dal.NwkAddr
	row := dalpsql.db.QueryRow("select nwkkey, nwkaddr from nwkaddrs left outer join sessions on sesnwk=nwkkey where sesactive is null or sesactive <= now() limit 1")
	err := row.Scan(&returnvalue.ID, &returnvalue.NwkAddr)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return &returnvalue, nil
}

func (dalpsql *dalPsql) SetActiveSessionsInactive(device int64) error {
	_, err := dalpsql.db.Exec("update sessions set sesactive = now() where sesdev=$1 and sesactive > now()")
	if err != nil {
		return err
	}
	return nil
}
