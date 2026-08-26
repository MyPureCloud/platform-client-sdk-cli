package telephony_prefixes_simulate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_prefixes_simulate_call"
)

func init() {
	telephony_prefixes_simulateCmd.AddCommand(telephony_prefixes_simulate_call.Cmdtelephony_prefixes_simulate_call())
	telephony_prefixes_simulateCmd.Short = utils.GenerateCustomDescription(telephony_prefixes_simulateCmd.Short, telephony_prefixes_simulate_call.Description, )
	telephony_prefixes_simulateCmd.Long = telephony_prefixes_simulateCmd.Short
}
