package main

func (p *Plugin) OnActivate() error {
	p.API.LogDebug("Activating plugin")

	p.API.LogDebug("Plugin activated")

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
