package integrations_botconnectors_outgoing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnectors_outgoing_messages"
)

func init() {
	integrations_botconnectors_outgoingCmd.AddCommand(integrations_botconnectors_outgoing_messages.Cmdintegrations_botconnectors_outgoing_messages())
	integrations_botconnectors_outgoingCmd.Short = utils.GenerateCustomDescription(integrations_botconnectors_outgoingCmd.Short, integrations_botconnectors_outgoing_messages.Description, )
	integrations_botconnectors_outgoingCmd.Long = integrations_botconnectors_outgoingCmd.Short
}
