package Engine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gamemann/Rust-Auto-Wipe/pkg/debug"
	"github.com/gamemann/tmc-servers-engine/internal/Config"
	"github.com/gamemann/tmc-servers-engine/pkg/TMCHttp"
)

func (e *QueryEngine) IPS4_UpdateServer(cfg Config.Config, results QueryResult, server Server) error {
	var err error

	id_str := strconv.FormatUint(uint64(server.ID), 10)

	// Compile URL we're going to send to.
	fullRequestURL := fmt.Sprintf("%s/%d", cfg.UpdateURL, server.ID)

	if cfg.BasicAuth {
		fullRequestURL = fmt.Sprintf("%s/%d?%s=%s", cfg.UpdateURL, server.ID, cfg.KeyParam, cfg.Token)
	}

	// Build headers.
	headers := make(map[string]string, 1)

	// If we're not using basic auth, set authorization header instead.
	if !cfg.BasicAuth {
		headers["Authorization"] = cfg.Token
	}

	// We're sending JSON ;)
	//headers["Content-Type"] = "application/json"

	// If our players max value is 0, it indicates the server is offline.
	if results.PlayersMax == nil || *results.PlayersMax < 1 {
		*results.Players = 0
		*results.MapName = ""
		*results.Users = nil
	}

	// Now build parameters.
	post_data := make(map[string]string, 1)

	// Check each one for nil values. If not nil, set POST data.
	if results.RealName != nil {
		post_data["realname"] = *results.RealName
	}

	if results.Players != nil {
		post_data["players"] = strconv.FormatUint(uint64(*results.Players), 10)
	}

	if results.PlayersMax != nil {
		post_data["playersmax"] = strconv.FormatUint(uint64(*results.PlayersMax), 10)
	}

	if results.MapName != nil {
		post_data["mapname"] = *results.MapName
	}

	if results.Verified != nil {
		post_data["verified"] = strconv.FormatUint(uint64(*results.Verified), 10)
	}

	// We want to update the last stat time as well!
	post_data["laststatupdate"] = strconv.FormatUint(uint64(time.Now().Unix()), 10)

	// Now send the request (POST for updating).
	data, rc, err := TMCHttp.SendHTTPReq(fullRequestURL, "POST", post_data, headers)

	if err != nil {
		debug.SendDebugMsg(id_str, int(cfg.Debug), 1, "Failed to update server.")

		return err
	}

	// Use debug package I made from another project :)
	debug.SendDebugMsg(id_str, int(cfg.Debug), 1, "Updated server "+server.IP+":"+strconv.FormatUint(uint64(server.Port), 10))
	debug.SendDebugMsg(id_str, int(cfg.Debug), 2, "Return Code => "+strconv.FormatUint(uint64(rc), 10))
	debug.SendDebugMsg(id_str, int(cfg.Debug), 2, "Body => "+data)

	return err
}

func (e *QueryEngine) IPS4_FetchServers(cfg Config.Config) error {
	var err error

	// Build headers.
	headers := make(map[string]string, 1)

	// If we're not using basic auth, set authorization header instead.
	if !cfg.BasicAuth {
		headers["Authorization"] = cfg.Token
	}

	// We're accepting and expecting JSON ;)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	// We're going to need to parse with page support.
	page := 1
	pages := 10

	// Clear current servers list.
	e.ServerList = []Server{}

	// Loop through pages.
	for page < pages {
		// Compile URL we're going to send to.
		fullRequestURL := fmt.Sprintf("%s?sort=%s&page=%d", cfg.UpdateURL, cfg.Sort, page)

		if cfg.BasicAuth {
			fullRequestURL = fmt.Sprintf("%s?sort=%s&key=%s&page=%d", cfg.UpdateURL, cfg.Sort, cfg.Token, page)
		}

		// Now send the request (POST for updating).
		data, rc, err := TMCHttp.SendHTTPReq(fullRequestURL, "GET", nil, headers)

		if err != nil {
			debug.SendDebugMsg("ALL", int(cfg.Debug), 2, "Error sending HTTP request => "+fullRequestURL)

			break
		}

		// Use debug package I made from another project :)
		debug.SendDebugMsg("ALL", int(cfg.Debug), 1, "Retrieving servers from IPS 4 API..")
		debug.SendDebugMsg("ALL", int(cfg.Debug), 2, "Return Code => "+strconv.FormatUint(uint64(rc), 10))
		debug.SendDebugMsg("ALL", int(cfg.Debug), 2, "Body => "+data)

		var resp Servers

		// Conversion from JSON.
		err = json.Unmarshal([]byte(data), &resp)

		if err != nil {
			debug.SendDebugMsg("ALL", int(cfg.Debug), 2, "Error parsing JSON data. Request => "+fullRequestURL)

			break
		}

		// Check if max pages needs to be changed.
		if pages != resp.TotalPages {
			pages = resp.TotalPages
		}

		// Now append all servers to the engine server list.
		e.ServerList = append(e.ServerList, resp.Results...)

		// Increment page count.
		page++

		// We're going to wait an interval to prevent rate limiting if possible.
		time.Sleep(time.Duration(cfg.WaitInterval) * time.Millisecond)
	}

	return err
}
