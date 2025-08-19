package integrations_botconnector

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnector_bots"
)

func init() {
	integrations_botconnectorCmd.AddCommand(integrations_botconnector_bots.Cmdintegrations_botconnector_bots())
	integrations_botconnectorCmd.Short = utils.GenerateCustomDescription(integrations_botconnectorCmd.Short, integrations_botconnector_bots.Description, )
	integrations_botconnectorCmd.Long = integrations_botconnectorCmd.Short
}
