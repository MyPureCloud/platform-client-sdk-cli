package integrations_actions_draft

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_validation"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_testfile"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_templates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft_function"
)

func init() {
	integrations_actions_draftCmd.AddCommand(integrations_actions_draft_validation.Cmdintegrations_actions_draft_validation())
	integrations_actions_draftCmd.AddCommand(integrations_actions_draft_testfile.Cmdintegrations_actions_draft_testfile())
	integrations_actions_draftCmd.AddCommand(integrations_actions_draft_schemas.Cmdintegrations_actions_draft_schemas())
	integrations_actions_draftCmd.AddCommand(integrations_actions_draft_templates.Cmdintegrations_actions_draft_templates())
	integrations_actions_draftCmd.AddCommand(integrations_actions_draft_publish.Cmdintegrations_actions_draft_publish())
	integrations_actions_draftCmd.AddCommand(integrations_actions_draft_function.Cmdintegrations_actions_draft_function())
	integrations_actions_draftCmd.Short = utils.GenerateCustomDescription(integrations_actions_draftCmd.Short, integrations_actions_draft_validation.Description, integrations_actions_draft_testfile.Description, integrations_actions_draft_schemas.Description, integrations_actions_draft_templates.Description, integrations_actions_draft_publish.Description, integrations_actions_draft_function.Description, )
	integrations_actions_draftCmd.Long = integrations_actions_draftCmd.Short
}
