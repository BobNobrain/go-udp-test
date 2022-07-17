package network

import (
	"errors"
	"goudptest/internal/util/binpack"
)

type Datagram interface {
	binpack.BinaryPackable
	GetTag() byte
}

func MarshallDatagram(bp *binpack.Binpacker, msg Datagram) {
	binpack.MarshallTagged(bp, msg.GetTag(), msg)
}

func UnmarshallDatagram(bp *binpack.Binpacker) (*Datagram, error) {
	return binpack.UnmarshallTagged(bp, createDatagram)
}

func createDatagram(tag byte) (Datagram, error) {
	switch tag {
	case CMTNoop:
		return &NoopCMsg{}, nil

	case CMTJoin:
		return &JoinCMsg{}, nil

	case CMTLeave:
		return &LeaveCMsg{}, nil

	case CMTMove:
		return &MoveCMsg{}, nil

	case SMTJoin:
		return &JoinSMsg{}, nil

	case SMTLeave:
		return &LeaveSMsg{}, nil

	case SMTKick:
		return &KickSMsg{}, nil

	case SMTUpdate:
		return &UpdateSMsg{}, nil
	}

	return nil, errors.New("unknown datagram tag")
}
