package integrations_actions_draft_function

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_function_upload"
)

func init() {
	integrations_actions_draft_functionCmd.AddCommand(integrations_actions_draft_function_upload.Cmdintegrations_actions_draft_function_upload())
	integrations_actions_draft_functionCmd.Short = utils.GenerateCustomDescription(integrations_actions_draft_functionCmd.Short, integrations_actions_draft_function_upload.Description, )
	integrations_actions_draft_functionCmd.Long = integrations_actions_draft_functionCmd.Short
}
