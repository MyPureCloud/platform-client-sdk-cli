package conversations_callbacks_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_callbacks_bulk_update"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_callbacks_bulk_disconnect"
)

func init() {
	conversations_callbacks_bulkCmd.AddCommand(conversations_callbacks_bulk_update.Cmdconversations_callbacks_bulk_update())
	conversations_callbacks_bulkCmd.AddCommand(conversations_callbacks_bulk_disconnect.Cmdconversations_callbacks_bulk_disconnect())
	conversations_callbacks_bulkCmd.Short = utils.GenerateCustomDescription(conversations_callbacks_bulkCmd.Short, conversations_callbacks_bulk_update.Description, conversations_callbacks_bulk_disconnect.Description, )
	conversations_callbacks_bulkCmd.Long = conversations_callbacks_bulkCmd.Short
}
