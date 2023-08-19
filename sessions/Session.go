package sessions

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"sukihcf.com/suki/database"
	"sukihcf.com/suki/sessions/rank"
	"time"
)

type Statistics struct {
	kills, killstreak, deaths, rankId int
}

var ScoreboardTicker = time.NewTicker(1 * time.Second)

type Session struct {
	Player *player.Player
	Stats  *Statistics
	Rank   *rank.Rank
}

func (s *Session) load() {
	var kills, deaths, killstreak, rankID int
	var username string

	data := database.GetPlayerData(s.Player.Name())
	if !data.Next() { //if we can't find the next row of data
		return
	}
	err := data.Scan(&username, &kills, &deaths, &killstreak, &rankID) //scan the data into the variables
	if err != nil {
		panic(err.Error())
		return
	}
	Stats := &Statistics{kills: kills, deaths: deaths, killstreak: killstreak, rankId: rankID}
	s.Stats = Stats
	s.Rank = rank.GetRankById(rankID)

	go s.CreateScoreboard()
}

func (s *Session) save() {
	kills := s.Kills()
	deaths := s.Deaths()
	killstreak := s.Killstreak()
	rankID := s.RankID()

	_, err := database.DB.SqlHandler.Exec("REPLACE INTO sukiPlayers (username, kills, deaths, killstreak, rankID) VALUES (?, ?, ?, ?, ?)", s.Player.Name(), kills, deaths, killstreak, rankID)
	if err != nil {
		panic(err.Error())
		return
	}

}

func (s *Session) CreateScoreboard() {
	for {
		select {
		case <-ScoreboardTicker.C:
			sb := scoreboard.New("Suki")
			s.Player.SendScoreboard(sb)
		}
	}
}

func (s *Session) DestroyScoreboard() {
	s.Player.RemoveScoreboard()
	ScoreboardTicker.Stop()
}

func (s *Session) SetRank(rank *rank.Rank) {
	s.Rank = rank
	s.Stats.rankId = rank.ID()
}

func (s *Session) GetRank() *rank.Rank {
	return s.Rank
}

func (s *Session) Kills() int {
	return s.Stats.kills
}

func (s *Session) Killstreak() int {
	return s.Stats.killstreak
}

func (s *Session) Deaths() int {
	return s.Stats.deaths
}

func (s *Session) RankID() int {
	return s.Stats.rankId
}
