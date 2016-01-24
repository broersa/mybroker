package bllimpl

import (
	"log"

	"github.com/broersa/ttnbroker/bll"
	"github.com/broersa/ttnbroker/dal"
)

type (
	bllImpl struct {
		dal dal.Dal
	}
)

// New Implemented Factory
func New(dal dal.Dal) bll.Bll {
	return &bllImpl{dal}
}

func (bllimpl *bllImpl) RegisterApplication(application *bll.Application) {
	log.Println("RegisterApplication")
}

func (bllimpl *bllImpl) ProcessJoinRequest(appeui string, deveui string, devnonce string) {
	//	session, err := bllimpl.dal.GetSession(devnonce)
}
