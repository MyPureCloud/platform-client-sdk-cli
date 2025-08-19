package outbound_campaigns_callback

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_campaigns_callback_schedule"
)

func init() {
	outbound_campaigns_callbackCmd.AddCommand(outbound_campaigns_callback_schedule.Cmdoutbound_campaigns_callback_schedule())
	outbound_campaigns_callbackCmd.Short = utils.GenerateCustomDescription(outbound_campaigns_callbackCmd.Short, outbound_campaigns_callback_schedule.Description, )
	outbound_campaigns_callbackCmd.Long = outbound_campaigns_callbackCmd.Short
}
