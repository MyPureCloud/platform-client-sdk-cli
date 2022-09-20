package integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_credentials"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_clientapps"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_types"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_userapps"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_botconnector"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech"
)

func init() {
	integrationsCmd.AddCommand(integrations_credentials.Cmdintegrations_credentials())
	integrationsCmd.AddCommand(integrations_clientapps.Cmdintegrations_clientapps())
	integrationsCmd.AddCommand(integrations_config.Cmdintegrations_config())
	integrationsCmd.AddCommand(integrations_types.Cmdintegrations_types())
	integrationsCmd.AddCommand(integrations_userapps.Cmdintegrations_userapps())
	integrationsCmd.AddCommand(integrations_actions.Cmdintegrations_actions())
	integrationsCmd.AddCommand(integrations_botconnector.Cmdintegrations_botconnector())
	integrationsCmd.AddCommand(integrations_speech.Cmdintegrations_speech())
	integrationsCmd.Short = utils.GenerateCustomDescription(integrationsCmd.Short, integrations_credentials.Description, integrations_clientapps.Description, integrations_config.Description, integrations_types.Description, integrations_userapps.Description, integrations_actions.Description, integrations_botconnector.Description, integrations_speech.Description, )
	integrationsCmd.Long = integrationsCmd.Short
}
