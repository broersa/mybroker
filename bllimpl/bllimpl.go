package bllimpl

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/broersa/mybroker/bll"
	"github.com/broersa/mybroker/dal"
)

type (
	bllImpl struct {
		dal *dal.Dal
	}
)

// New Implemented Factory
func New(dal *dal.Dal) bll.Bll {
	return &bllImpl{dal}
}

func (bllimpl *bllImpl) RegisterApplication(appname string) (int64, error) {
	a, err := (*bllimpl.dal).GetApplicationOnName(appname)
	if err != nil {
		return 0, err
	}
	if a != nil {
		return 0, errors.New("Application allready exists")
	}

	appeui := make([]byte, 8)
	_, err = rand.Read(appeui)
	if err != nil {
		return 0, err
	}

	str := hex.EncodeToString(appeui)

	application := &dal.Application{Name: appname, AppEUI: str}
	(*bllimpl.dal).BeginTransaction()
	id, err := (*bllimpl.dal).AddApplication(application)
	if err != nil {
		return 0, err
	}
	(*bllimpl.dal).CommitTransaction()
	return id, nil
}

func (bllimpl *bllImpl) GetApplication(id int64) (*bll.Application, error) {
	a, err := (*bllimpl.dal).GetApplication(id)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, nil
	}
	return &bll.Application{ID: a.ID, Name: a.Name, AppEUI: a.AppEUI}, nil
}

func (bllimpl *bllImpl) GetApplicationOnAppEUI(appeui string) (*bll.Application, error) {
	a, err := (*bllimpl.dal).GetApplicationOnAppEUI(appeui)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, nil
	}
	return &bll.Application{ID: a.ID, Name: a.Name, AppEUI: a.AppEUI}, nil
}

func (bllimpl *bllImpl) RegisterDevice(appeui string, deveui string) (int64, error) {
	d, err := (*bllimpl.dal).GetDeviceOnDevEUI(deveui)
	if err != nil {
		return 0, err
	}
	if d != nil {
		return 0, errors.New("Device allready exists")
	}

	a, err := (*bllimpl.dal).GetApplicationOnAppEUI(appeui)
	if err != nil {
		return 0, err
	}
	if a == nil {
		return 0, errors.New("Application does not exists")
	}

	appkey := make([]byte, 16)
	_, err = rand.Read(appkey)
	if err != nil {
		return 0, err
	}

	str := hex.EncodeToString(appkey)

	device := &dal.Device{Application: a.ID, DevEUI: deveui, AppKey: str}
	(*bllimpl.dal).BeginTransaction()
	id, err := (*bllimpl.dal).AddDevice(device)
	if err != nil {
		return 0, err
	}
	(*bllimpl.dal).CommitTransaction()
	return id, nil
}

func (bllimpl *bllImpl) GetDevice(id int64) (*bll.Device, error) {
	d, err := (*bllimpl.dal).GetDevice(id)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, nil
	}
	return &bll.Device{ID: d.ID, Application: d.Application, DevEUI: d.DevEUI, AppKey: d.AppKey}, nil
}

func (bllimpl *bllImpl) GetDeviceOnAppEUIDevEUI(appeui string, deveui string) (*bll.Device, error) {
	d, err := (*bllimpl.dal).GetDeviceOnAppEUIDevEUI(appeui, deveui)
	if err != nil {
		return nil, err
	}
	if d == nil {
		return nil, nil
	}
	return &bll.Device{ID: d.ID, Application: d.Application, DevEUI: d.DevEUI, AppKey: d.AppKey}, nil
}

func (bllimpl *bllImpl) getAppNonce() ([]byte, error) {
	returnvalue := make([]byte, 3)
	_, err := rand.Read(returnvalue)
	if err != nil {
		return nil, err
	}
	return returnvalue, nil
}

func (bllimpl *bllImpl) ProcessJoinRequest(appeui string, deveui string, devnonce string) (uint32, []byte, error) {
	(*bllimpl.dal).BeginTransaction()
	// check if device exists
	device, err := (*bllimpl.dal).GetDeviceOnAppEUIDevEUI(appeui, deveui)
	if err != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, err
	}
	if device == nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, errors.New("AppEUI, DevEUI does not exists")
	}
	// check if there is allready an active session with the devnonce
	session, err := (*bllimpl.dal).GetSessionOnDeviceDevNonceActive(device.ID, devnonce)
	if err != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, err
	}
	if session != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, errors.New("Session allready active")
	}
	nwkaddr, err := (*bllimpl.dal).GetFreeNwkAddr()
	if err != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, err
	}
	appnonce, err := bllimpl.getAppNonce()
	if err != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, err
	}

	err = (*bllimpl.dal).SetActiveSessionsInactive(device.ID)
	if err != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, err
	}

	appnoncestr := hex.EncodeToString(appnonce)
	_, err = (*bllimpl.dal).AddSession(&dal.Session{Device: device.ID, NwkAddr: nwkaddr.ID, DevNonce: devnonce, AppNonce: appnoncestr, NwkSKey: "", AppSKey: ""})
	if err != nil {
		(*bllimpl.dal).RollbackTransaction()
		return 0, nil, err
	}
	//	session, err := bllimpl.dal.GetSession(devnonce)
	(*bllimpl.dal).CommitTransaction()
	return nwkaddr.NwkAddr, appnonce, nil
}
