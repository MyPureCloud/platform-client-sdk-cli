package outbound_dnclists

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_uploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_customexclusioncolumns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_emailaddresses"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_export"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_importstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_phonenumbers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_dnclists_whatsappnumbers"
)

func init() {
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_uploads.Cmdoutbound_dnclists_uploads())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_customexclusioncolumns.Cmdoutbound_dnclists_customexclusioncolumns())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_divisionviews.Cmdoutbound_dnclists_divisionviews())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_emailaddresses.Cmdoutbound_dnclists_emailaddresses())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_export.Cmdoutbound_dnclists_export())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_importstatus.Cmdoutbound_dnclists_importstatus())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_phonenumbers.Cmdoutbound_dnclists_phonenumbers())
	outbound_dnclistsCmd.AddCommand(outbound_dnclists_whatsappnumbers.Cmdoutbound_dnclists_whatsappnumbers())
	outbound_dnclistsCmd.Short = utils.GenerateCustomDescription(outbound_dnclistsCmd.Short, outbound_dnclists_uploads.Description, outbound_dnclists_customexclusioncolumns.Description, outbound_dnclists_divisionviews.Description, outbound_dnclists_emailaddresses.Description, outbound_dnclists_export.Description, outbound_dnclists_importstatus.Description, outbound_dnclists_phonenumbers.Description, outbound_dnclists_whatsappnumbers.Description, )
	outbound_dnclistsCmd.Long = outbound_dnclistsCmd.Short
}
