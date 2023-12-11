package outbound_importtemplates_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_importtemplates_bulk_add"
)

func init() {
	outbound_importtemplates_bulkCmd.AddCommand(outbound_importtemplates_bulk_add.Cmdoutbound_importtemplates_bulk_add())
	outbound_importtemplates_bulkCmd.Short = utils.GenerateCustomDescription(outbound_importtemplates_bulkCmd.Short, outbound_importtemplates_bulk_add.Description, )
	outbound_importtemplates_bulkCmd.Long = outbound_importtemplates_bulkCmd.Short
}
