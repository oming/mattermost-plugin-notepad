package main

import (
	"encoding/json"
)

type Bookmark struct {
	ChannelID       string `json:"channelId"`
	BookmarkContent string `json:"bookmark_content"`
}

func (p *Plugin) GetBookmark(channelID string) (*Bookmark, error) {
	p.API.LogDebug("GetBookmark Start. channelID: " + channelID)
	bookmarkBytes, appErr := p.API.KVGet(channelID)
	if appErr != nil {
		p.API.LogError("KVGet Error. channelID: " + channelID)
		return nil, appErr
	}

	var bookmark *Bookmark
	if bookmarkBytes != nil {
		p.API.LogDebug("bookmarkBytes != nil")
		if err := json.Unmarshal(bookmarkBytes, &bookmark); err != nil {
			p.API.LogDebug("bookmarkBytes error")
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
		bookmark = &Bookmark{
			ChannelID:       channelID,
			BookmarkContent: "내용을 등록하세요.",
		}
	}

	return bookmark, nil
}

func (p *Plugin) SaveBookmark(bookmark *Bookmark) error {
	p.API.LogDebug("SaveBookmark Start.", bookmark)
	jsonBookmark, err := json.Marshal(bookmark)
	if err != nil {
		return err
	}

	if appErr := p.API.KVSet(bookmark.ChannelID, jsonBookmark); appErr != nil {
		return appErr
	}

	return nil
}
