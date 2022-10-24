package Engine

import (
	"fmt"

	"github.com/rumblefrog/go-a2s"
)

// The ID of the engine stored in the TMC database (1 = A2S Engine).
const ID = 1

func (e *Engine) A2S_Query(server Server) (QueryResult, error) {
	var result QueryResult
	var general_info *a2s.ServerInfo
	var err error

	var realname string
	var players uint
	var playersmax uint
	var mapname string

	conn_str := fmt.Sprintf("%s:%d", server.IP, int(server.Port))

	// Use A2S package to send query with challenge support!
	a2s_query, err := a2s.NewClient(conn_str, a2s.SetMaxPacketSize(14000))

	// If we have an error, just go to the end of the function to return.
	if err != nil {
		goto end
	}

	// Close connection at end.
	defer a2s_query.Close()

	// Retrieve general information.
	general_info, err = a2s_query.QueryInfo()

	if err != nil || general_info == nil {
		goto end
	}

	// Copy result variables and cast to what we need.
	realname = general_info.Name
	players = uint(general_info.Players)
	playersmax = uint(general_info.MaxPlayers)
	mapname = general_info.Map

	// Map values.
	result.RealName = &realname
	result.Players = &players
	result.PlayersMax = &playersmax
	result.MapName = &mapname

end:
	return result, err
}
