package Engine

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gamemann/Rust-Auto-Wipe/pkg/debug"
	"github.com/gamemann/tmc-servers-engine/internal/Config"
	"github.com/gamemann/tmc-servers-engine/pkg/TMCHttp"
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

var lastupdate map[int]int

func (e *QueryEngine) Handler(cfg *Config.Config) {
	lastupdate = make(map[int]int)
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
			now := time.Now().Unix()

			// Make queries to retrieve up-to-date server stats/info.
			if e.ClassName == "A2S" {
				qr, err = e.A2S_Query(srv)

				if err != nil {
					debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 2, "Failed to send A2S_INFO request to "+srv.IP+":"+strconv.FormatUint(uint64(srv.Port), 10))
					debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 2, err.Error())

					continue
				}

				// Now check if we should verify the server.
				if qr.RealName != nil && srv.ClaimKey != nil && strings.Contains(*qr.RealName, *srv.ClaimKey) {
					one := uint(1)
					qr.Verified = &one
				}
			}

			// Send an update request if needed.
			if e.APIName == "IPS4" {
				e.IPS4_UpdateServer(*cfg, qr, srv)
			}

			// Post hook.
			if lastupdate[srv.ID] < 1 || uint(now) > uint(lastupdate[srv.ID])+cfg.PostHookInterval {

				fullRequestURL := fmt.Sprintf("%s/%d", cfg.PostHook, srv.ID)

				if cfg.BasicAuth {
					fullRequestURL = fmt.Sprintf("%s/%d?key=%s", cfg.PostHook, srv.ID, cfg.Token)
				}

				// Send POST request (e.g. for stats updating).
				d, rc, err := TMCHttp.SendHTTPReq(fullRequestURL, "GET", nil, nil, false)

				if err != nil {
					debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 1, "Failed to send post hook request.")
				} else {
					debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 2, "Sent post hook (GET request).")
				}

				debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 2, "Request URL => "+fullRequestURL)
				debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 2, "Return Code => "+strconv.FormatUint(uint64(rc), 10))
				debug.SendDebugMsg(strconv.FormatUint(uint64(srv.ID), 10), int(cfg.Debug), 3, "Return Body => "+d)

				lastupdate[srv.ID] = int(now)
			}

			// Wait time.
			time.Sleep(time.Duration(cfg.WaitInterval) * time.Millisecond)
		}

		// Fetch time.
		time.Sleep(time.Duration(cfg.FetchInterval) * time.Millisecond)
	}
}
