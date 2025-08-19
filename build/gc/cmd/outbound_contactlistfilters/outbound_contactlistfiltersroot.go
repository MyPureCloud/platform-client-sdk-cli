package outbound_contactlistfilters

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlistfilters_preview"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlistfilters_bulk"
)

func init() {
	outbound_contactlistfiltersCmd.AddCommand(outbound_contactlistfilters_preview.Cmdoutbound_contactlistfilters_preview())
	outbound_contactlistfiltersCmd.AddCommand(outbound_contactlistfilters_bulk.Cmdoutbound_contactlistfilters_bulk())
	outbound_contactlistfiltersCmd.Short = utils.GenerateCustomDescription(outbound_contactlistfiltersCmd.Short, outbound_contactlistfilters_preview.Description, outbound_contactlistfilters_bulk.Description, )
	outbound_contactlistfiltersCmd.Long = outbound_contactlistfiltersCmd.Short
}
