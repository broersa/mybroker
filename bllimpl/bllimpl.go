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

func (bllimpl *bllImpl) ProcessJoinRequest(appeui string, deveui string, devnonce string) {
	//	session, err := bllimpl.dal.GetSession(devnonce)
}
