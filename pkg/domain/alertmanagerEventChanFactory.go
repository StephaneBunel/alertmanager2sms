package domain

import (
	"github.com/romana/rlog"
)

func NewEventChan(qlen int) AmEventChan {
	if qlen < 1 {
		qlen = 128
		rlog.Info("set alertmanger event queue channel to", qlen, "because it is lower than 1.")
	}
	return make(AmEventChan, qlen)
}
