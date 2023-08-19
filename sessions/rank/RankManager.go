package rank

import "github.com/df-mc/dragonfly/server"

var ranks = make(map[string]*Rank)

func addRank(rank *Rank) {
	ranks[rank.Name()] = rank
}

func RegisterRanks() {
	addRank(&Rank{name: "Default", id: 0, format: "§7%name%: %message%"})
	addRank(&Rank{name: "VIP", id: 1, format: "§a[VIP] %name%: %message%"})
}

func GetRankById(id int) *Rank {
	for _, rank := range ranks {
		if rank.ID() == id {
			return rank
		}
	}
	return nil
}

func BroadcastMessage(message string, server *server.Server) {
	for _, p := range server.Players() {
		p.Message(message)
	}
}

func GetDefaultRank() *Rank {
	return GetRankById(0)
}

func GetRank(name string) *Rank {
	return ranks[name]
}

func GetRanks() map[string]*Rank {
	return ranks
}
