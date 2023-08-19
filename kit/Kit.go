package kit

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

type Kit interface {
	Armor() []item.Stack
	Items() []item.Stack
	Apply(p *player.Player)
	Effects() []effect.Effect
}

func ApplyKit(kit *Kit, p *player.Player) {
	(*kit).Apply(p)
}
