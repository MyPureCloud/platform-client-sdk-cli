package outbound_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules_campaigns"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_schedules_sequences"
)

func init() {
	outbound_schedulesCmd.AddCommand(outbound_schedules_campaigns.Cmdoutbound_schedules_campaigns())
	outbound_schedulesCmd.AddCommand(outbound_schedules_sequences.Cmdoutbound_schedules_sequences())
	outbound_schedulesCmd.Short = utils.GenerateCustomDescription(outbound_schedulesCmd.Short, outbound_schedules_campaigns.Description, outbound_schedules_sequences.Description, )
	outbound_schedulesCmd.Long = outbound_schedulesCmd.Short
}
