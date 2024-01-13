package integrations_actions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_categories"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_certificates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_draft"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_testfile"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_templates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_execute"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions_drafts"
)

func init() {
	integrations_actionsCmd.AddCommand(integrations_actions_categories.Cmdintegrations_actions_categories())
	integrations_actionsCmd.AddCommand(integrations_actions_certificates.Cmdintegrations_actions_certificates())
	integrations_actionsCmd.AddCommand(integrations_actions_draft.Cmdintegrations_actions_draft())
	integrations_actionsCmd.AddCommand(integrations_actions_testfile.Cmdintegrations_actions_testfile())
	integrations_actionsCmd.AddCommand(integrations_actions_schemas.Cmdintegrations_actions_schemas())
	integrations_actionsCmd.AddCommand(integrations_actions_templates.Cmdintegrations_actions_templates())
	integrations_actionsCmd.AddCommand(integrations_actions_execute.Cmdintegrations_actions_execute())
	integrations_actionsCmd.AddCommand(integrations_actions_drafts.Cmdintegrations_actions_drafts())
	integrations_actionsCmd.Short = utils.GenerateCustomDescription(integrations_actionsCmd.Short, integrations_actions_categories.Description, integrations_actions_certificates.Description, integrations_actions_draft.Description, integrations_actions_testfile.Description, integrations_actions_schemas.Description, integrations_actions_templates.Description, integrations_actions_execute.Description, integrations_actions_drafts.Description, )
	integrations_actionsCmd.Long = integrations_actionsCmd.Short
}
