package integrations_unifiedcommunications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_unifiedcommunications_thirdpartypresences"
)

func init() {
	integrations_unifiedcommunicationsCmd.AddCommand(integrations_unifiedcommunications_thirdpartypresences.Cmdintegrations_unifiedcommunications_thirdpartypresences())
	integrations_unifiedcommunicationsCmd.Short = utils.GenerateCustomDescription(integrations_unifiedcommunicationsCmd.Short, integrations_unifiedcommunications_thirdpartypresences.Description, )
	integrations_unifiedcommunicationsCmd.Long = integrations_unifiedcommunicationsCmd.Short
}
