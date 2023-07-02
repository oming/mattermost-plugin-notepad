package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/pkg/errors"
)

func (p *Plugin) setupBotUser() error {
	botUserID, err := p.API.EnsureBotUser(&model.Bot{
		Username:    "notepad",
		DisplayName: "Notepad",
		Description: "Bot for notepad plugin.",
	})
	if err != nil {
		p.API.LogError("Error in setting up bot user", "Error", err.Error())
		return err
	}

	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		return err
	}

	profileImage, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", "icon.png"))
	if err != nil {
		return err
	}

	if appErr := p.API.SetProfileImage(botUserID, profileImage); appErr != nil {
		return errors.Wrap(appErr, "couldn't set profile image")
	}

	p.BotUserID = botUserID
	return nil
}

func (p *Plugin) OnActivate() error {
	p.API.LogDebug("Activating plugin")

	p.API.LogDebug("Plugin activated")
	if err := p.setupBotUser(); err != nil {
		p.API.LogError("Failed to create a bot user", "Error", err.Error())
		return err
	}

	return p.API.RegisterCommand(getCommand())
}

func (p *Plugin) OnDeactivate() error {
	if p.telemetryClient != nil {
		err := p.telemetryClient.Close()
		if err != nil {
			p.API.LogWarn("OnDeactivate: Failed to close telemetryClient", "error", err.Error())
		}
	}
	return nil
}
