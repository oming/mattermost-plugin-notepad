package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/google/go-cmp/cmp"
	"github.com/mattermost/mattermost-plugin-api/experimental/telemetry"
	"github.com/mattermost/mattermost-server/v6/model"
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

	BotUserID string
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
	case "/notepad":
		p.httpNotepadSettings(w, r)
	default:
		http.NotFound(w, r)
	}
}
func (p *Plugin) httpNotepadSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("httpNotepadSettings Start")
	switch r.Method {
	case http.MethodPost:
		p.httpNotepadSaveSettings(w, r)
	case http.MethodGet:
		p.httpNotepadGetSettings(w, r)
	default:
		http.Error(w, "Request: "+r.Method+" is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (p *Plugin) httpNotepadSaveSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("httpNotepadSaveSettings Start")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var notepad *Notepad
	if err = json.Unmarshal(body, &notepad); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orgNotepad, err2 := p.GetNotepad(notepad.ChannelID)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	if err = p.SaveNotepad(notepad); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.Header.Get("Mattermost-User-Id")
	mmuser, _ := p.API.GetUser(userID)
	username := mmuser.Username

	if diff := cmp.Diff(orgNotepad.NotepadContent, notepad.NotepadContent); diff != "" {
		msg := "#### Notepad has been updated.\n"
		msg += "Modifier @" + username + " Diff ( '-' = removed, '+' = added )\n\n"
		msg += "---\n"
		msg += "```\n" +
			"%s" +
			"\n```"
		_, err3 := p.API.CreatePost(&model.Post{
			UserId:    p.BotUserID,
			ChannelId: notepad.ChannelID,
			Message:   fmt.Sprintf(msg, diff),
		})
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}
	}

	resp := struct {
		Status string
	}{"OK"}

	p.writeJSON(w, resp)
}

func (p *Plugin) httpNotepadGetSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("httpNotepadGetSettings Start")
	channelID, ok := r.URL.Query()["channelId"]

	if !ok || len(channelID[0]) < 1 {
		http.Error(w, "Missing channelId parameter", http.StatusBadRequest)
		return
	}
	p.API.LogDebug("Channel ID " + channelID[0])

	notepad, err := p.GetNotepad(channelID[0])
	p.API.LogDebug("notepad 값 확인", notepad)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := ResponseNotepad{
		ChannelID:      notepad.ChannelID,
		ChannelNotepad: notepad.NotepadContent,
		CommonNotepad:  p.configuration.CommonNotepad,
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

type ResponseNotepad struct {
	ChannelID      string `json:"channelId"`
	ChannelNotepad string `json:"channelNotepad"`
	CommonNotepad  string `json:"commonNotepad"`
}
