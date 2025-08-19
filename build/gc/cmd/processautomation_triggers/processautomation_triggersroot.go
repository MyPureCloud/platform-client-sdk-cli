package processautomation_triggers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/processautomation_triggers_topics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/processautomation_triggers_testfile"
)

func init() {
	processautomation_triggersCmd.AddCommand(processautomation_triggers_topics.Cmdprocessautomation_triggers_topics())
	processautomation_triggersCmd.AddCommand(processautomation_triggers_testfile.Cmdprocessautomation_triggers_testfile())
	processautomation_triggersCmd.Short = utils.GenerateCustomDescription(processautomation_triggersCmd.Short, processautomation_triggers_topics.Description, processautomation_triggers_testfile.Description, )
	processautomation_triggersCmd.Long = processautomation_triggersCmd.Short
}
