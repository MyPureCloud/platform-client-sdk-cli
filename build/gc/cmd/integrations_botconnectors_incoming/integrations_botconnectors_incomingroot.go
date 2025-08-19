package integrations_botconnectors_incoming

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnectors_incoming_messages"
)

func init() {
	integrations_botconnectors_incomingCmd.AddCommand(integrations_botconnectors_incoming_messages.Cmdintegrations_botconnectors_incoming_messages())
	integrations_botconnectors_incomingCmd.Short = utils.GenerateCustomDescription(integrations_botconnectors_incomingCmd.Short, integrations_botconnectors_incoming_messages.Description, )
	integrations_botconnectors_incomingCmd.Long = integrations_botconnectors_incomingCmd.Short
}
