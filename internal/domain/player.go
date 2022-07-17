package domain

import (
	"goudptest/internal/util/binpack"
	"goudptest/internal/util/geom"
)

type Player struct {
	Name     string
	Position geom.Vec3
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

func (player *Player) MoveTo(newPos geom.Vec3) {
	player.Position = newPos
}

func (player *Player) Marshall(bp *binpack.Binpacker) {
	bp.MarshallString16(player.Name)
	player.Position.Marshall(bp)
}
func (player *Player) Unmarshall(bp *binpack.Binpacker) error {
	name, nerr := bp.UnmarshallString16()
	if nerr != nil {
		return nerr
	}

	player.Name = name

	perr := player.Position.Unmarshall(bp)
	if perr != nil {
		return perr
	}

	return nil
}
