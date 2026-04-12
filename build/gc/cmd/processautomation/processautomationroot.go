package processautomation

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/processautomation_triggers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/processautomation_scheduledtriggers"
)

func init() {
	processautomationCmd.AddCommand(processautomation_triggers.Cmdprocessautomation_triggers())
	processautomationCmd.AddCommand(processautomation_scheduledtriggers.Cmdprocessautomation_scheduledtriggers())
	processautomationCmd.Short = utils.GenerateCustomDescription(processautomationCmd.Short, processautomation_triggers.Description, processautomation_scheduledtriggers.Description, )
	processautomationCmd.Long = processautomationCmd.Short
}
