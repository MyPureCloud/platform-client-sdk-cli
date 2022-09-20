package processautomation

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/processautomation_triggers"
)

func init() {
	processautomationCmd.AddCommand(processautomation_triggers.Cmdprocessautomation_triggers())
	processautomationCmd.Short = utils.GenerateCustomDescription(processautomationCmd.Short, processautomation_triggers.Description, )
	processautomationCmd.Long = processautomationCmd.Short
}
