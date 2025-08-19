package processautomation_triggers_topics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/processautomation_triggers_topics_testfile"
)

func init() {
	processautomation_triggers_topicsCmd.AddCommand(processautomation_triggers_topics_testfile.Cmdprocessautomation_triggers_topics_testfile())
	processautomation_triggers_topicsCmd.Short = utils.GenerateCustomDescription(processautomation_triggers_topicsCmd.Short, processautomation_triggers_topics_testfile.Description, )
	processautomation_triggers_topicsCmd.Long = processautomation_triggers_topicsCmd.Short
}
