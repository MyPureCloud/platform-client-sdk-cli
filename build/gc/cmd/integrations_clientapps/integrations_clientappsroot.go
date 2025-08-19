package integrations_clientapps

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_clientapps_unifiedcommunications"
)

func init() {
	integrations_clientappsCmd.AddCommand(integrations_clientapps_unifiedcommunications.Cmdintegrations_clientapps_unifiedcommunications())
	integrations_clientappsCmd.Short = utils.GenerateCustomDescription(integrations_clientappsCmd.Short, integrations_clientapps_unifiedcommunications.Description, )
	integrations_clientappsCmd.Long = integrations_clientappsCmd.Short
}
