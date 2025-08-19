package integrations_actions_functions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_functions_runtimes"
)

func init() {
	integrations_actions_functionsCmd.AddCommand(integrations_actions_functions_runtimes.Cmdintegrations_actions_functions_runtimes())
	integrations_actions_functionsCmd.Short = utils.GenerateCustomDescription(integrations_actions_functionsCmd.Short, integrations_actions_functions_runtimes.Description, )
	integrations_actions_functionsCmd.Long = integrations_actions_functionsCmd.Short
}
