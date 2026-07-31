package agentic

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/agentic_virtualagents"
)

func init() {
	agenticCmd.AddCommand(agentic_virtualagents.Cmdagentic_virtualagents())
	agenticCmd.Short = utils.GenerateCustomDescription(agenticCmd.Short, agentic_virtualagents.Description, )
	agenticCmd.Long = agenticCmd.Short
}
