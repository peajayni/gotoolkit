package ids

import (
	"github.com/google/uuid"
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

func NewUUIDs() *UUIDs {
	return &UUIDs{}
}

type UUIDs struct{}

func (u *UUIDs) New() string {
	guid := uuid.New()
	return guid.String()
}
