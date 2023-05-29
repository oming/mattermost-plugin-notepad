package main

import (
	"encoding/json"
	"io"
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
	// tracker         telemetry.Tracker
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("ServeHTTP Start")
	w.Header().Set("Content-Type", "application/json")

	userID := r.Header.Get("Mattermost-User-Id")
	if userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	p.API.LogDebug("UserID: " + userID)

	switch path := r.URL.Path; path {
	case "/bookmark":
		p.httpBookmarkSettings(w, r)
	default:
		http.NotFound(w, r)
	}
}
func (p *Plugin) httpBookmarkSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("httpBookmarkSettings Start")
	switch r.Method {
	case http.MethodPost:
		p.httpBookmarkSaveSettings(w, r)
	case http.MethodGet:
		p.httpBookmarkGetSettings(w, r)
	default:
		http.Error(w, "Request: "+r.Method+" is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (p *Plugin) httpBookmarkSaveSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("httpBookmarkSaveSettings Start")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var bookmark *Bookmark
	if err = json.Unmarshal(body, &bookmark); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = p.SaveBookmark(bookmark); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		Status string
	}{"OK"}

	p.writeJSON(w, resp)
}

func (p *Plugin) httpBookmarkGetSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("httpBookmarkGetSettings Start")
	channelID, ok := r.URL.Query()["channelId"]

	if !ok || len(channelID[0]) < 1 {
		http.Error(w, "Missing channelId parameter", http.StatusBadRequest)
		return
	}
	p.API.LogDebug("Channel ID " + channelID[0])

	bookmark, err := p.GetBookmark(channelID[0])
	p.API.LogDebug("bookmark 값 확인", bookmark)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := ResponseBookmark{
		ChannelID:       bookmark.ChannelID,
		ChannelBookmark: bookmark.BookmarkContent,
		CommonBookmark:  p.configuration.CommonBookmark,
	}

	p.writeJSON(w, resp)
}

func (p *Plugin) handleErrorWithCode(w http.ResponseWriter, code int, errTitle string, err error) {
	w.WriteHeader(code)
	b, _ := json.Marshal(struct {
		Error   string `json:"error"`
		Details string `json:"details"`
	}{
		Error:   errTitle,
		Details: err.Error(),
	})
	_, _ = w.Write(b)
}

func (p *Plugin) writeJSON(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		p.API.LogWarn("Failed to marshal JSON response", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		p.API.LogWarn("Failed to write JSON response", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/

type ResponseBookmark struct {
	ChannelID       string `json:"channelId"`
	ChannelBookmark string `json:"channelBookmark"`
	CommonBookmark  string `json:"commonBookmark"`
}
