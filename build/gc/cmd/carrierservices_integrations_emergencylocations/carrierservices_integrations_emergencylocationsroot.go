package carrierservices_integrations_emergencylocations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/carrierservices_integrations_emergencylocations_me"
)

func init() {
	carrierservices_integrations_emergencylocationsCmd.AddCommand(carrierservices_integrations_emergencylocations_me.Cmdcarrierservices_integrations_emergencylocations_me())
	carrierservices_integrations_emergencylocationsCmd.Short = utils.GenerateCustomDescription(carrierservices_integrations_emergencylocationsCmd.Short, carrierservices_integrations_emergencylocations_me.Description, )
	carrierservices_integrations_emergencylocationsCmd.Long = carrierservices_integrations_emergencylocationsCmd.Short
}
