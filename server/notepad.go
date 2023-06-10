package main

import (
	"encoding/json"
)

type Notepad struct {
	ChannelID      string `json:"channelId"`
	NotepadContent string `json:"notepad_content"`
}

func (p *Plugin) GetNotepad(channelID string) (*Notepad, error) {
	p.API.LogDebug("GetNotepad Start. channelID: " + channelID)
	notepadBytes, appErr := p.API.KVGet(channelID)
	if appErr != nil {
		p.API.LogError("KVGet Error. channelID: " + channelID)
		return nil, appErr
	}

	var notepad *Notepad
	if notepadBytes != nil {
		p.API.LogDebug("notepadBytes != nil")
		if err := json.Unmarshal(notepadBytes, &notepad); err != nil {
			p.API.LogDebug("notepadBytes error")
			return nil, err
		}
	} else {
		p.API.LogDebug("Return Default Values")
		// Return a default value
		channel, err := p.API.GetChannel(channelID)
		if err != nil {
			return nil, err
		}
		p.API.LogDebug("channel: ", channel.Name)
		notepad = &Notepad{
			ChannelID:      channelID,
			NotepadContent: "내용을 등록하세요.",
		}
	}

	return notepad, nil
}

func (p *Plugin) SaveNotepad(notepad *Notepad) error {
	p.API.LogDebug("SaveNotepad Start.", notepad)
	jsonNotepad, err := json.Marshal(notepad)
	if err != nil {
		return err
	}

	if appErr := p.API.KVSet(notepad.ChannelID, jsonNotepad); appErr != nil {
		return appErr
	}

	return nil
}
