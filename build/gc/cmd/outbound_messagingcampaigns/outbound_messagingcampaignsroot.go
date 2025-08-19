package outbound_messagingcampaigns

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_messagingcampaigns_progress"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_messagingcampaigns_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_messagingcampaigns_diagnostics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_messagingcampaigns_start"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_messagingcampaigns_stop"
)

func init() {
	outbound_messagingcampaignsCmd.AddCommand(outbound_messagingcampaigns_progress.Cmdoutbound_messagingcampaigns_progress())
	outbound_messagingcampaignsCmd.AddCommand(outbound_messagingcampaigns_divisionviews.Cmdoutbound_messagingcampaigns_divisionviews())
	outbound_messagingcampaignsCmd.AddCommand(outbound_messagingcampaigns_diagnostics.Cmdoutbound_messagingcampaigns_diagnostics())
	outbound_messagingcampaignsCmd.AddCommand(outbound_messagingcampaigns_start.Cmdoutbound_messagingcampaigns_start())
	outbound_messagingcampaignsCmd.AddCommand(outbound_messagingcampaigns_stop.Cmdoutbound_messagingcampaigns_stop())
	outbound_messagingcampaignsCmd.Short = utils.GenerateCustomDescription(outbound_messagingcampaignsCmd.Short, outbound_messagingcampaigns_progress.Description, outbound_messagingcampaigns_divisionviews.Description, outbound_messagingcampaigns_diagnostics.Description, outbound_messagingcampaigns_start.Description, outbound_messagingcampaigns_stop.Description, )
	outbound_messagingcampaignsCmd.Long = outbound_messagingcampaignsCmd.Short
}
