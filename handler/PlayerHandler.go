package handler

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"strings"
	"sukihcf.com/suki/sessions"
	"sukihcf.com/suki/sessions/rank"
)

type PlayerHandler struct {
	player.NopHandler //means playerHandler embeds aka implements player.NopHandler
	server            *server.Server
	player            *player.Player
}

func (h *PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	p := h.player
	session := sessions.GetSession(p)

	playerRank := session.GetRank()
	replacer := strings.NewReplacer("%name%", p.Name(), "%message%", *message)

	replaced := replacer.Replace(playerRank.Format())
	ctx.Cancel()

	*message = replaced
	rank.BroadcastMessage(*message, h.server)

}

func (h *PlayerHandler) HandleQuit() {
	session := sessions.GetSession(h.player)
	if session == nil {
		return
	}

	sessions.RemoveSession(h.player)
}

func CreateHandler(player *player.Player, server *server.Server) *PlayerHandler {
	return &PlayerHandler{
		player: player,
		server: server,
	}
}
