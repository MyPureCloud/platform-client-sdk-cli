package architect_prompts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_prompts_resources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_prompts_history"
)

func init() {
	architect_promptsCmd.AddCommand(architect_prompts_resources.Cmdarchitect_prompts_resources())
	architect_promptsCmd.AddCommand(architect_prompts_history.Cmdarchitect_prompts_history())
	architect_promptsCmd.Short = utils.GenerateCustomDescription(architect_promptsCmd.Short, architect_prompts_resources.Description, architect_prompts_history.Description, )
	architect_promptsCmd.Long = architect_promptsCmd.Short
}
