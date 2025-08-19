package integrations_botconnector_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnector_bots_summaries"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnector_bots_versions"
)

func init() {
	integrations_botconnector_botsCmd.AddCommand(integrations_botconnector_bots_summaries.Cmdintegrations_botconnector_bots_summaries())
	integrations_botconnector_botsCmd.AddCommand(integrations_botconnector_bots_versions.Cmdintegrations_botconnector_bots_versions())
	integrations_botconnector_botsCmd.Short = utils.GenerateCustomDescription(integrations_botconnector_botsCmd.Short, integrations_botconnector_bots_summaries.Description, integrations_botconnector_bots_versions.Description, )
	integrations_botconnector_botsCmd.Long = integrations_botconnector_botsCmd.Short
}
