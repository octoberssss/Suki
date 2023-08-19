package sessions

import (
	"github.com/df-mc/dragonfly/server/player"
	"sync"
)

var sessions = make(map[string]*Session)
var mtx sync.Mutex

// SessionFactory is a factory for creating sessions.

func Sessions() map[string]*Session {
	defer mtx.Unlock()
	mtx.Lock()
	return sessions
}

func CreateSession(player *player.Player) {
	mtx.Lock()
	defer mtx.Unlock()
	sessions[player.Name()] = &Session{Player: player, Stats: &Statistics{}}
	session := sessions[player.Name()]

	session.load()
}

func RemoveSession(player *player.Player) {
	mtx.Lock()
	defer mtx.Unlock()
	session := sessions[player.Name()]

	session.DestroyScoreboard()
	session.save()
	delete(sessions, player.Name())
}

func hasSession(player *player.Player) bool {
	defer mtx.Unlock()
	mtx.Lock()

	return sessions[player.Name()] != nil
}

func GetSession(p *player.Player) *Session {
	name := p.Name()
	return sessions[name]
}
