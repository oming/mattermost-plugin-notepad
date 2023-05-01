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
	p.API.LogDebug("hsan", "trigger", trigger)

	switch trigger {
	case "show":
		return p.executeCommandShow(args), nil
	default:
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         fmt.Sprintf("Unknown command: " + args.Command),
		}, nil
	}
}
func (p *Plugin) executeCommandShow(args *model.CommandArgs) *model.CommandResponse {
	serverConfig := p.API.GetConfig()
	configuration := p.getConfiguration()
	p.API.LogDebug("hsan", "args", args)
	p.API.LogDebug("hsan", "serverConfig", serverConfig)
	p.API.LogDebug("hsan", "configuration", configuration)

	_ = p.API.SendEphemeralPost(args.UserId, &model.Post{
		ChannelId: args.ChannelId,
		Message:   configuration.BookmarkContent,
		// Props: model.StringInterface{
		// 	"type": "custom_demo_plugin_ephemeral",
		// },
	})
	return &model.CommandResponse{}
}

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "bookmark",
		DisplayName:      "Bookmark Bot",
		Description:      "Interact with your Todo list.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: show",
		AutoCompleteHint: "[command]",
		AutocompleteData: getAutocompleteData(),
	}
}

func getAutocompleteData() *model.AutocompleteData {
	bookmark := model.NewAutocompleteData("bookmark", "[command]", "Available commands: show")
	show := model.NewAutocompleteData("show", "", "Display Bookmark")
	bookmark.AddCommand(show)
	return bookmark
}
