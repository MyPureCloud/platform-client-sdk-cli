package outbound_filespecificationtemplates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_filespecificationtemplates_bulk"
)

func init() {
	outbound_filespecificationtemplatesCmd.AddCommand(outbound_filespecificationtemplates_bulk.Cmdoutbound_filespecificationtemplates_bulk())
	outbound_filespecificationtemplatesCmd.Short = utils.GenerateCustomDescription(outbound_filespecificationtemplatesCmd.Short, outbound_filespecificationtemplates_bulk.Description, )
	outbound_filespecificationtemplatesCmd.Long = outbound_filespecificationtemplatesCmd.Short
}
