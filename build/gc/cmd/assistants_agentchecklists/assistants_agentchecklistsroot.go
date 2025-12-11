package assistants_agentchecklists

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_agentchecklists_languages"
)

func init() {
	assistants_agentchecklistsCmd.AddCommand(assistants_agentchecklists_languages.Cmdassistants_agentchecklists_languages())
	assistants_agentchecklistsCmd.Short = utils.GenerateCustomDescription(assistants_agentchecklistsCmd.Short, assistants_agentchecklists_languages.Description, )
	assistants_agentchecklistsCmd.Long = assistants_agentchecklistsCmd.Short
}
