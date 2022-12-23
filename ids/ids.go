package ids

import (
	"github.com/rs/xid"
)

func NewXids() *Xids {
	return &Xids{}
}

type Xids struct{}

func (x *Xids) New() string {
	guid := xid.New()
	return guid.String()
}
