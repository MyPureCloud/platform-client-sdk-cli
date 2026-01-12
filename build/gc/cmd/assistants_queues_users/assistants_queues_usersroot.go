package assistants_queues_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues_users_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues_users_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues_users_query"
)

func init() {
	assistants_queues_usersCmd.AddCommand(assistants_queues_users_jobs.Cmdassistants_queues_users_jobs())
	assistants_queues_usersCmd.AddCommand(assistants_queues_users_bulk.Cmdassistants_queues_users_bulk())
	assistants_queues_usersCmd.AddCommand(assistants_queues_users_query.Cmdassistants_queues_users_query())
	assistants_queues_usersCmd.Short = utils.GenerateCustomDescription(assistants_queues_usersCmd.Short, assistants_queues_users_jobs.Description, assistants_queues_users_bulk.Description, assistants_queues_users_query.Description, )
	assistants_queues_usersCmd.Long = assistants_queues_usersCmd.Short
}
