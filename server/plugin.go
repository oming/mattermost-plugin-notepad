package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-plugin-api/experimental/telemetry"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	telemetryClient telemetry.Client
	tracker         telemetry.Tracker
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("hsan", "ServeHTTP", c.RequestId)

	userID := r.Header.Get("Mattermost-User-Id")
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method == http.MethodGet {
		if r.URL.Path == "/bookmark" {
			p.handleBookmark(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func (p *Plugin) handleBookmark(w http.ResponseWriter, _ *http.Request) {
	info := map[string]interface{}{
		"bookmark": p.getConfiguration().BookmarkContent,
	}
	p.API.LogDebug("hsan", "handleBookmark", info)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(info); err != nil {
		p.API.LogError(err.Error())
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
