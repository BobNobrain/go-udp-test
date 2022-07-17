package network

import (
	"goudptest/internal/util/binpack"
	"goudptest/internal/util/geom"
)

// Messages that can be sent from client
type ClientMsgType = byte

const (
	CMTNoop  ClientMsgType = 0
	CMTJoin  ClientMsgType = 1
	CMTLeave ClientMsgType = 2
	CMTMove  ClientMsgType = 3
)

type NoopCMsg struct{}

func (*NoopCMsg) GetTag() byte { return CMTNoop }
func (*NoopCMsg) Marshall(bp *binpack.Binpacker) {
}
func (msg *NoopCMsg) Unmarshall(bp *binpack.Binpacker) error {
	return nil
}

type JoinCMsg struct {
	PlayerName string
}

func (*JoinCMsg) GetTag() byte { return CMTJoin }
func (msg JoinCMsg) Marshall(bp *binpack.Binpacker) {
	bp.MarshallString16(msg.PlayerName)
}
func (msg *JoinCMsg) Unmarshall(bp *binpack.Binpacker) error {
	name, err := bp.UnmarshallString16()
	msg.PlayerName = name
	return err
}

type LeaveCMsg struct {
}

func (*LeaveCMsg) GetTag() byte { return CMTLeave }
func (LeaveCMsg) Marshall(bp *binpack.Binpacker) {
}
func (msg *LeaveCMsg) Unmarshall(bp *binpack.Binpacker) error {
	return nil
}

type MoveCMsg struct {
	NewPosition geom.Vec3
}

func (*MoveCMsg) GetTag() byte { return CMTMove }
func (msg MoveCMsg) Marshall(bp *binpack.Binpacker) {
	msg.NewPosition.Marshall(bp)
}
func (msg *MoveCMsg) Unmarshall(bp *binpack.Binpacker) error {
	return msg.NewPosition.Unmarshall(bp)
}
