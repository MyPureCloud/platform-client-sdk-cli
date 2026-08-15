package intents_customerintents_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_customerintents_bulk_retrieve"
)

func init() {
	intents_customerintents_bulkCmd.AddCommand(intents_customerintents_bulk_retrieve.Cmdintents_customerintents_bulk_retrieve())
	intents_customerintents_bulkCmd.Short = utils.GenerateCustomDescription(intents_customerintents_bulkCmd.Short, intents_customerintents_bulk_retrieve.Description, )
	intents_customerintents_bulkCmd.Long = intents_customerintents_bulkCmd.Short
}
