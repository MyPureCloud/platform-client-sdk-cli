package integrations_credentials

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_credentials_listing"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_credentials_types"
)

func init() {
	integrations_credentialsCmd.AddCommand(integrations_credentials_listing.Cmdintegrations_credentials_listing())
	integrations_credentialsCmd.AddCommand(integrations_credentials_types.Cmdintegrations_credentials_types())
	integrations_credentialsCmd.Short = utils.GenerateCustomDescription(integrations_credentialsCmd.Short, integrations_credentials_listing.Description, integrations_credentials_types.Description, )
	integrations_credentialsCmd.Long = integrations_credentialsCmd.Short
}
