package integrations_botconnectors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnectors_bots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnectors_incoming"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnectors_outgoing"
)

func init() {
	integrations_botconnectorsCmd.AddCommand(integrations_botconnectors_bots.Cmdintegrations_botconnectors_bots())
	integrations_botconnectorsCmd.AddCommand(integrations_botconnectors_incoming.Cmdintegrations_botconnectors_incoming())
	integrations_botconnectorsCmd.AddCommand(integrations_botconnectors_outgoing.Cmdintegrations_botconnectors_outgoing())
	integrations_botconnectorsCmd.Short = utils.GenerateCustomDescription(integrations_botconnectorsCmd.Short, integrations_botconnectors_bots.Description, integrations_botconnectors_incoming.Description, integrations_botconnectors_outgoing.Description, )
	integrations_botconnectorsCmd.Long = integrations_botconnectorsCmd.Short
}
