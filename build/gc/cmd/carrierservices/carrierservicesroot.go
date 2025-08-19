package carrierservices

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/carrierservices_integrations"
)

func init() {
	carrierservicesCmd.AddCommand(carrierservices_integrations.Cmdcarrierservices_integrations())
	carrierservicesCmd.Short = utils.GenerateCustomDescription(carrierservicesCmd.Short, carrierservices_integrations.Description, )
	carrierservicesCmd.Long = carrierservicesCmd.Short
}
