package assistants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_agentchecklists"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_copilot"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues"
)

func init() {
	assistantsCmd.AddCommand(assistants_agentchecklists.Cmdassistants_agentchecklists())
	assistantsCmd.AddCommand(assistants_copilot.Cmdassistants_copilot())
	assistantsCmd.AddCommand(assistants_queues.Cmdassistants_queues())
	assistantsCmd.Short = utils.GenerateCustomDescription(assistantsCmd.Short, assistants_agentchecklists.Description, assistants_copilot.Description, assistants_queues.Description, )
	assistantsCmd.Long = assistantsCmd.Short
}
