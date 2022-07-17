package domain

import (
	"errors"
	"goudptest/internal/util/geom"
	"sync"
)

type World struct {
	players        map[string]*Player
	playerLimit    int
	playerCount    int
	updatedPlayers map[string]*Player
	// TODO: more optimal mutex usage?
	lock sync.RWMutex
}

func NewWorld(playerLimit int) *World {
	return &World{
		players:        make(map[string]*Player),
		playerLimit:    playerLimit,
		playerCount:    0,
		updatedPlayers: make(map[string]*Player),
	}
}

func (world *World) RegisterPlayer(name string) error {
	world.lock.Lock()
	defer world.lock.Unlock()

	if world.playerCount >= world.playerLimit {
		return errors.New("this world is full")
	}

	world.players[name] = NewPlayer(name)
	world.playerCount += 1
	return nil
}

func (world *World) UnregisterPlayer(name string) error {
	world.lock.Lock()
	defer world.lock.Unlock()

	_, ok := world.players[name]
	if !ok {
		return errors.New("this player is not registered")
	}

	delete(world.players, name)
	world.playerCount -= 1
	return nil
}

func (world *World) MovePlayerTo(name string, newPos geom.Vec3) error {
	world.lock.Lock()
	defer world.lock.Unlock()

	p, ok := world.players[name]
	if !ok {
		return errors.New("no such player")
	}

	if newPos.Minus(p.Position).SizeSquared() > 25 {
		return errors.New("you're moving too fast")
	}

	p.MoveTo(newPos)

	world.updatedPlayers[p.Name] = p

	return nil
}
func (world *World) MovePlayerBy(name string, dPos geom.Vec3) (geom.Vec3, error) {
	world.lock.RLock()
	newPos := world.players[name].Position.Plus(dPos)
	world.lock.RUnlock()

	return newPos, world.MovePlayerTo(name, newPos)
}

func (world *World) GatherPlayerUpdates() []*Player {
	world.lock.Lock()
	defer world.lock.Unlock()

	result := make([]*Player, len(world.updatedPlayers))

	i := 0
	for _, p := range world.updatedPlayers {
		copy := *p
		result[i] = &copy
		i += 1
	}

	world.updatedPlayers = make(map[string]*Player)

	return result
}

func (world *World) ApplyPlayerUpdates(updates []*Player) error {
	world.lock.Lock()
	defer world.lock.Unlock()

	for _, u := range updates {
		p, ok := world.players[u.Name]
		if !ok {
			return errors.New("no such player")
		}

		p.Position = u.Position
	}

	return nil
}
