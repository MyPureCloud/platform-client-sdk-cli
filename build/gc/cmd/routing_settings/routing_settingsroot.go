package routing_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_settings_transcription"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_settings_contactcenter"
)

func init() {
	routing_settingsCmd.AddCommand(routing_settings_transcription.Cmdrouting_settings_transcription())
	routing_settingsCmd.AddCommand(routing_settings_contactcenter.Cmdrouting_settings_contactcenter())
	routing_settingsCmd.Short = utils.GenerateCustomDescription(routing_settingsCmd.Short, routing_settings_transcription.Description, routing_settings_contactcenter.Description, )
	routing_settingsCmd.Long = routing_settingsCmd.Short
}
