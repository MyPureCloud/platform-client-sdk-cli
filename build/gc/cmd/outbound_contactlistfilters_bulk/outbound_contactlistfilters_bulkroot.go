package outbound_contactlistfilters_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlistfilters_bulk_retrieve"
)

func init() {
	outbound_contactlistfilters_bulkCmd.AddCommand(outbound_contactlistfilters_bulk_retrieve.Cmdoutbound_contactlistfilters_bulk_retrieve())
	outbound_contactlistfilters_bulkCmd.Short = utils.GenerateCustomDescription(outbound_contactlistfilters_bulkCmd.Short, outbound_contactlistfilters_bulk_retrieve.Description, )
	outbound_contactlistfilters_bulkCmd.Long = outbound_contactlistfilters_bulkCmd.Short
}
