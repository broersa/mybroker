package bllimpl

import (
	"log"

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

func (bllimpl *bllImpl) RegisterApplication(application *bll.Application) {
	log.Println("RegisterApplication")
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
