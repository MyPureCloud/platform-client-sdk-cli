package outbound_contactlisttemplates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlisttemplates_bulk"
)

func init() {
	outbound_contactlisttemplatesCmd.AddCommand(outbound_contactlisttemplates_bulk.Cmdoutbound_contactlisttemplates_bulk())
	outbound_contactlisttemplatesCmd.Short = utils.GenerateCustomDescription(outbound_contactlisttemplatesCmd.Short, outbound_contactlisttemplates_bulk.Description, )
	outbound_contactlisttemplatesCmd.Long = outbound_contactlisttemplatesCmd.Short
}
