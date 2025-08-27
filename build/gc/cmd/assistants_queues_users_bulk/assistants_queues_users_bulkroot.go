package assistants_queues_users_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues_users_bulk_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/assistants_queues_users_bulk_remove"
)

func init() {
	assistants_queues_users_bulkCmd.AddCommand(assistants_queues_users_bulk_add.Cmdassistants_queues_users_bulk_add())
	assistants_queues_users_bulkCmd.AddCommand(assistants_queues_users_bulk_remove.Cmdassistants_queues_users_bulk_remove())
	assistants_queues_users_bulkCmd.Short = utils.GenerateCustomDescription(assistants_queues_users_bulkCmd.Short, assistants_queues_users_bulk_add.Description, assistants_queues_users_bulk_remove.Description, )
	assistants_queues_users_bulkCmd.Long = assistants_queues_users_bulkCmd.Short
}
