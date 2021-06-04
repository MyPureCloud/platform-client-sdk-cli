package integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_actions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_credentials"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_eventlog"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_types"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_clientapps"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_workforcemanagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_userapps"
)

func init() {
	integrationsCmd.AddCommand(integrations_speech.Cmdintegrations_speech())
	integrationsCmd.AddCommand(integrations_config.Cmdintegrations_config())
	integrationsCmd.AddCommand(integrations_actions.Cmdintegrations_actions())
	integrationsCmd.AddCommand(integrations_credentials.Cmdintegrations_credentials())
	integrationsCmd.AddCommand(integrations_eventlog.Cmdintegrations_eventlog())
	integrationsCmd.AddCommand(integrations_types.Cmdintegrations_types())
	integrationsCmd.AddCommand(integrations_clientapps.Cmdintegrations_clientapps())
	integrationsCmd.AddCommand(integrations_workforcemanagement.Cmdintegrations_workforcemanagement())
	integrationsCmd.AddCommand(integrations_userapps.Cmdintegrations_userapps())
	integrationsCmd.Short = utils.GenerateCustomDescription(integrationsCmd.Short, integrations_speech.Description, integrations_config.Description, integrations_actions.Description, integrations_credentials.Description, integrations_eventlog.Description, integrations_types.Description, integrations_clientapps.Description, integrations_workforcemanagement.Description, integrations_userapps.Description, )
	integrationsCmd.Long = integrationsCmd.Short
}
