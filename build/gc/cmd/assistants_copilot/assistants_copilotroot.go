package assistants_copilot

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_copilot_featuresupport"
)

func init() {
	assistants_copilotCmd.AddCommand(assistants_copilot_featuresupport.Cmdassistants_copilot_featuresupport())
	assistants_copilotCmd.Short = utils.GenerateCustomDescription(assistants_copilotCmd.Short, assistants_copilot_featuresupport.Description, )
	assistants_copilotCmd.Long = assistants_copilotCmd.Short
}
