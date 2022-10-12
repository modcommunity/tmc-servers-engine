package Engine

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gamemann/Rust-Auto-Wipe/pkg/debug"
	"github.com/gamemann/tmc-servers-engine/internal/Config"
)

type QueryEngine struct {
	ClassName  string
	APIName    string
	ServerList []Server
}

type QueryResult struct {
	RealName   *string `json:"realname"`
	Users      *[]User `json:"users"`
	Players    *uint   `json:"players"`
	PlayersMax *uint   `json:"playersmax"`
	MapName    *string `json:"mapname"`

	// These are for verification.
	Verified *uint `json:"verified"`
	NoVerify *uint `json:"noverify"`
}

func (e *QueryEngine) Handler(cfg *Config.Config) {
	// Create a loop since this is another thread.
	for {
		// Fetch servers.
		if e.APIName == "IPS4" {
			e.IPS4_FetchServers(*cfg)
		}

		// Loop through each server.
		for _, srv := range e.ServerList {
			var qr QueryResult
			var err error

			// Make queries to retrieve up-to-date server stats/info.
			if e.ClassName == "A2S" {

				qr, err = e.A2S_Query(srv)

				if err != nil {
					debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 2, "Failed to send A2S_INFO request to "+srv.IP+":"+strconv.FormatUint(uint64(srv.Port), 10))
					fmt.Println(err)

					continue
				}
			}

			// Send an update request if needed.
			if e.APIName == "IPS4" {
				e.IPS4_UpdateServer(*cfg, qr, srv)
			}

			// Wait time.
			time.Sleep(time.Duration(cfg.WaitInterval))
		}

		// Fetch time.
		time.Sleep(time.Duration(cfg.FetchInterval))
	}
}
