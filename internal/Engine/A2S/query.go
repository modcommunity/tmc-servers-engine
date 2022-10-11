package A2S

import (
	"fmt"

	"github.com/gamemann/tmc-servers-engine/internal/Engine"
	"github.com/rumblefrog/go-a2s"
)

// The ID of the engine stored in the TMC database (1 = A2S Engine).
const ID = 1

func Query(ip string, port uint) (Engine.QueryResult, error) {
	var result Engine.QueryResult
	var general_info *a2s.ServerInfo
	var err error

	conn_str := fmt.Sprintf("%s:%d", ip, int(port))

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

	// Map values.
	result.RealName = general_info.Name
	result.Players = uint(general_info.Players)
	result.PlayersMax = uint(general_info.MaxPlayers)
	result.MapName = general_info.Map

end:
	return result, err
}
