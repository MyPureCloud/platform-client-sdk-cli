package assistants_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues_users"
)

func init() {
	assistants_queuesCmd.AddCommand(assistants_queues_users.Cmdassistants_queues_users())
	assistants_queuesCmd.Short = utils.GenerateCustomDescription(assistants_queuesCmd.Short, assistants_queues_users.Description, )
	assistants_queuesCmd.Long = assistants_queuesCmd.Short
}
