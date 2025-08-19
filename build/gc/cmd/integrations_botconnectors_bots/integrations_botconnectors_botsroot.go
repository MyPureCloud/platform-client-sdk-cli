package integrations_botconnectors_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnectors_bots_summaries"
)

func init() {
	integrations_botconnectors_botsCmd.AddCommand(integrations_botconnectors_bots_summaries.Cmdintegrations_botconnectors_bots_summaries())
	integrations_botconnectors_botsCmd.Short = utils.GenerateCustomDescription(integrations_botconnectors_botsCmd.Short, integrations_botconnectors_bots_summaries.Description, )
	integrations_botconnectors_botsCmd.Long = integrations_botconnectors_botsCmd.Short
}
