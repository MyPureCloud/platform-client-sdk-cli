package outbound_contactlists

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_contacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_importstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_timezonemappingpreview"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_clear"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_contactlists_export"
)

func init() {
	outbound_contactlistsCmd.AddCommand(outbound_contactlists_contacts.Cmdoutbound_contactlists_contacts())
	outbound_contactlistsCmd.AddCommand(outbound_contactlists_importstatus.Cmdoutbound_contactlists_importstatus())
	outbound_contactlistsCmd.AddCommand(outbound_contactlists_timezonemappingpreview.Cmdoutbound_contactlists_timezonemappingpreview())
	outbound_contactlistsCmd.AddCommand(outbound_contactlists_divisionviews.Cmdoutbound_contactlists_divisionviews())
	outbound_contactlistsCmd.AddCommand(outbound_contactlists_clear.Cmdoutbound_contactlists_clear())
	outbound_contactlistsCmd.AddCommand(outbound_contactlists_export.Cmdoutbound_contactlists_export())
	outbound_contactlistsCmd.Short = utils.GenerateCustomDescription(outbound_contactlistsCmd.Short, outbound_contactlists_contacts.Description, outbound_contactlists_importstatus.Description, outbound_contactlists_timezonemappingpreview.Description, outbound_contactlists_divisionviews.Description, outbound_contactlists_clear.Description, outbound_contactlists_export.Description, )
	outbound_contactlistsCmd.Long = outbound_contactlistsCmd.Short
}
