package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

// ExecuteCommand executes a given command and returns a command response.
func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	trigger := strings.TrimPrefix(strings.Fields(args.Command)[1], "/")

	switch trigger {
	case "channel":
		return p.executeCommandChannel(args), nil
	case "common":
		return p.executeCommandCommon(args), nil
	default:
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         fmt.Sprintf("Unknown command: " + args.Command),
		}, nil
	}
}
func (p *Plugin) executeCommandChannel(args *model.CommandArgs) *model.CommandResponse {
	notepad, _ := p.GetNotepad(args.ChannelId)

	_ = p.API.SendEphemeralPost(args.UserId, &model.Post{
		ChannelId: args.ChannelId,
		Message:   notepad.NotepadContent,
	})
	return &model.CommandResponse{}
}
func (p *Plugin) executeCommandCommon(args *model.CommandArgs) *model.CommandResponse {
	configuration := p.getConfiguration()

	_ = p.API.SendEphemeralPost(args.UserId, &model.Post{
		ChannelId: args.ChannelId,
		Message:   configuration.CommonNotepad,
	})
	return &model.CommandResponse{}
}

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "notepad",
		DisplayName:      "Notepad Bot",
		Description:      "Show Notepad.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: channel, common",
		AutoCompleteHint: "[command]",
		AutocompleteData: getAutocompleteData(),
	}
}

func getAutocompleteData() *model.AutocompleteData {
	notepad := model.NewAutocompleteData("notepad", "[command]", "Available commands: channel, common")
	channel := model.NewAutocompleteData("channel", "", "Display Channel Notepad")
	notepad.AddCommand(channel)
	common := model.NewAutocompleteData("common", "", "Display Common Notepad")
	notepad.AddCommand(common)
	return notepad
}
