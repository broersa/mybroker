package models

import (
	"github.com/broersa/semtech"
)

// ResponseMessage Data Entity
type ResponseMessage struct {
	OriginUDPAddrNetwork string       `json:"originudpaddrnetwork"`
	OriginUDPAddrString  string       `json:"originudpaddrstring"`
	Package              semtech.TXPK `json:"package"`
}
