package carrierservices_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/carrierservices_integrations_emergencylocations"
)

func init() {
	carrierservices_integrationsCmd.AddCommand(carrierservices_integrations_emergencylocations.Cmdcarrierservices_integrations_emergencylocations())
	carrierservices_integrationsCmd.Short = utils.GenerateCustomDescription(carrierservices_integrationsCmd.Short, carrierservices_integrations_emergencylocations.Description, )
	carrierservices_integrationsCmd.Long = carrierservices_integrationsCmd.Short
}
