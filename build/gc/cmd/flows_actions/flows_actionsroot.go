package flows_actions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions_revert"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions_checkin"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions_checkout"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions_deactivate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_actions_unlock"
)

func init() {
	flows_actionsCmd.AddCommand(flows_actions_publish.Cmdflows_actions_publish())
	flows_actionsCmd.AddCommand(flows_actions_revert.Cmdflows_actions_revert())
	flows_actionsCmd.AddCommand(flows_actions_checkin.Cmdflows_actions_checkin())
	flows_actionsCmd.AddCommand(flows_actions_checkout.Cmdflows_actions_checkout())
	flows_actionsCmd.AddCommand(flows_actions_deactivate.Cmdflows_actions_deactivate())
	flows_actionsCmd.AddCommand(flows_actions_unlock.Cmdflows_actions_unlock())
	flows_actionsCmd.Short = utils.GenerateCustomDescription(flows_actionsCmd.Short, flows_actions_publish.Description, flows_actions_revert.Description, flows_actions_checkin.Description, flows_actions_checkout.Description, flows_actions_deactivate.Description, flows_actions_unlock.Description, )
	flows_actionsCmd.Long = flows_actionsCmd.Short
}
