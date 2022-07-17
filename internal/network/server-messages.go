package network

import (
	"goudptest/internal/domain"
	"goudptest/internal/util/binpack"
)

// Messages that can be sent from server
type ServerMsgType = byte

const (
	SMTJoin   ServerMsgType = 128
	SMTLeave  ServerMsgType = 129
	SMTKick   ServerMsgType = 130
	SMTUpdate ServerMsgType = 131
)

type JoinSMsg struct {
	PlayerName string
}

func (*JoinSMsg) GetTag() byte { return SMTJoin }
func (msg *JoinSMsg) Marshall(bp *binpack.Binpacker) {
	bp.MarshallString16(msg.PlayerName)
}
func (msg *JoinSMsg) Unmarshall(bp *binpack.Binpacker) error {
	name, err := bp.UnmarshallString16()
	msg.PlayerName = name
	return err
}

type LeaveSMsg struct {
	PlayerName string
}

func (*LeaveSMsg) GetTag() byte { return SMTLeave }
func (msg *LeaveSMsg) Marshall(bp *binpack.Binpacker) {
	bp.MarshallString16(msg.PlayerName)
}
func (msg *LeaveSMsg) Unmarshall(bp *binpack.Binpacker) error {
	name, err := bp.UnmarshallString16()
	msg.PlayerName = name
	return err
}

type KickSMsg struct {
	Reason string
}

func (*KickSMsg) GetTag() byte { return SMTKick }
func (msg *KickSMsg) Marshall(bp *binpack.Binpacker) {
	bp.MarshallString16(msg.Reason)
}
func (msg *KickSMsg) Unmarshall(bp *binpack.Binpacker) error {
	name, err := bp.UnmarshallString16()
	msg.Reason = name
	return err
}

type UpdateSMsg struct {
	UpdatedPlayers []*domain.Player
}

func (*UpdateSMsg) GetTag() byte { return SMTUpdate }
func (msg *UpdateSMsg) Marshall(bp *binpack.Binpacker) {
	binpack.MarshallArray16(bp, msg.UpdatedPlayers)
}
func (msg *UpdateSMsg) Unmarshall(bp *binpack.Binpacker) error {
	updates, err := binpack.UnmarshallArray16[*domain.Player](bp)
	msg.UpdatedPlayers = updates
	return err
}
