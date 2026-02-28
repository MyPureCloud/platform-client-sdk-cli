package telephony_calls

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_calls_metrics"
)

func init() {
	telephony_callsCmd.AddCommand(telephony_calls_metrics.Cmdtelephony_calls_metrics())
	telephony_callsCmd.Short = utils.GenerateCustomDescription(telephony_callsCmd.Short, telephony_calls_metrics.Description, )
	telephony_callsCmd.Long = telephony_callsCmd.Short
}
