package outbound_contactlists_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_bulk_update"
)

func init() {
	outbound_contactlists_bulkCmd.AddCommand(outbound_contactlists_bulk_update.Cmdoutbound_contactlists_bulk_update())
	outbound_contactlists_bulkCmd.Short = utils.GenerateCustomDescription(outbound_contactlists_bulkCmd.Short, outbound_contactlists_bulk_update.Description, )
	outbound_contactlists_bulkCmd.Long = outbound_contactlists_bulkCmd.Short
}
