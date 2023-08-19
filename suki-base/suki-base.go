package suki_base

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/pelletier/go-toml"
	"os"
	"sukihcf.com/suki/database"
	"sukihcf.com/suki/handler"
	"sukihcf.com/suki/sessions/rank"
)

type SukiBase struct {
	Name       string
	MaxPlayers int
	Server     *server.Server
}

var WorldManager *handler.WorldManager

func (suki *SukiBase) getName() string {
	return suki.Name
}

func (suki *SukiBase) getMaxPlayers() int {
	return suki.MaxPlayers
}

func (suki *SukiBase) Start() {
	s := suki.Server

	s.CloseOnProgramEnd()
	s.Listen()

	worldHandler := handler.CreateWorldManager(suki.Server, "worlds")
	WorldManager = worldHandler

	database.OpenDB()
	rank.RegisterRanks()

}

func GetWorldManager() *handler.WorldManager {
	return WorldManager
}

func ReadConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
