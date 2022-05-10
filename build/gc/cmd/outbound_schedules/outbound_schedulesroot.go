package outbound_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules_campaigns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules_sequences"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules_emailcampaigns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules_messagingcampaigns"
)

func init() {
	outbound_schedulesCmd.AddCommand(outbound_schedules_campaigns.Cmdoutbound_schedules_campaigns())
	outbound_schedulesCmd.AddCommand(outbound_schedules_sequences.Cmdoutbound_schedules_sequences())
	outbound_schedulesCmd.AddCommand(outbound_schedules_emailcampaigns.Cmdoutbound_schedules_emailcampaigns())
	outbound_schedulesCmd.AddCommand(outbound_schedules_messagingcampaigns.Cmdoutbound_schedules_messagingcampaigns())
	outbound_schedulesCmd.Short = utils.GenerateCustomDescription(outbound_schedulesCmd.Short, outbound_schedules_campaigns.Description, outbound_schedules_sequences.Description, outbound_schedules_emailcampaigns.Description, outbound_schedules_messagingcampaigns.Description, )
	outbound_schedulesCmd.Long = outbound_schedulesCmd.Short
}
