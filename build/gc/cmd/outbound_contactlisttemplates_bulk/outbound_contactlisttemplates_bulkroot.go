package outbound_contactlisttemplates_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlisttemplates_bulk_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlisttemplates_bulk_retrieve"
)

func init() {
	outbound_contactlisttemplates_bulkCmd.AddCommand(outbound_contactlisttemplates_bulk_add.Cmdoutbound_contactlisttemplates_bulk_add())
	outbound_contactlisttemplates_bulkCmd.AddCommand(outbound_contactlisttemplates_bulk_retrieve.Cmdoutbound_contactlisttemplates_bulk_retrieve())
	outbound_contactlisttemplates_bulkCmd.Short = utils.GenerateCustomDescription(outbound_contactlisttemplates_bulkCmd.Short, outbound_contactlisttemplates_bulk_add.Description, outbound_contactlisttemplates_bulk_retrieve.Description, )
	outbound_contactlisttemplates_bulkCmd.Long = outbound_contactlisttemplates_bulkCmd.Short
}
