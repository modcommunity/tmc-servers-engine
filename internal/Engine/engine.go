package Engine

import (
	"github.com/gamemann/tmc-servers-engine/internal/Engine/A2S"
	"github.com/gamemann/tmc-servers-engine/internal/Server"
)

type Engine struct {
	ClassName  string
	ServerList []Server.Server
}

type QueryResult struct {
	RealName   string        `json:"realname"`
	PlayerList []Server.User `json:"users"`
	Players    uint          `json:"players"`
	PlayersMax uint          `json:"playersmax"`
	MapName    string        `json:"mapname"`
}

// Technically we can support any protocol (e.g. UDP or TCP) :)
func (e Engine) MakeQuery(ip string, port uint) QueryResult {
	var query QueryResult

	// We must check the classname.
	if e.ClassName == "A2S" {
		query = A2S.Query(ip, port)
	}

	return query

}
