package outbound_importtemplates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_importtemplates_importstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_importtemplates_bulk"
)

func init() {
	outbound_importtemplatesCmd.AddCommand(outbound_importtemplates_importstatus.Cmdoutbound_importtemplates_importstatus())
	outbound_importtemplatesCmd.AddCommand(outbound_importtemplates_bulk.Cmdoutbound_importtemplates_bulk())
	outbound_importtemplatesCmd.Short = utils.GenerateCustomDescription(outbound_importtemplatesCmd.Short, outbound_importtemplates_importstatus.Description, outbound_importtemplates_bulk.Description, )
	outbound_importtemplatesCmd.Long = outbound_importtemplatesCmd.Short
}
