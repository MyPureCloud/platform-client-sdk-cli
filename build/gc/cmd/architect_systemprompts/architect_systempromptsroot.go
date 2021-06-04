package architect_systemprompts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_systemprompts_resources"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_systemprompts_history"
)

func init() {
	architect_systempromptsCmd.AddCommand(architect_systemprompts_resources.Cmdarchitect_systemprompts_resources())
	architect_systempromptsCmd.AddCommand(architect_systemprompts_history.Cmdarchitect_systemprompts_history())
	architect_systempromptsCmd.Short = utils.GenerateCustomDescription(architect_systempromptsCmd.Short, architect_systemprompts_resources.Description, architect_systemprompts_history.Description, )
	architect_systempromptsCmd.Long = architect_systempromptsCmd.Short
}
