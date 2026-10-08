package telephony_prefixes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_prefixes_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_prefixes_simulate"
)

func init() {
	telephony_prefixesCmd.AddCommand(telephony_prefixes_bulk.Cmdtelephony_prefixes_bulk())
	telephony_prefixesCmd.AddCommand(telephony_prefixes_simulate.Cmdtelephony_prefixes_simulate())
	telephony_prefixesCmd.Short = utils.GenerateCustomDescription(telephony_prefixesCmd.Short, telephony_prefixes_bulk.Description, telephony_prefixes_simulate.Description, )
	telephony_prefixesCmd.Long = telephony_prefixesCmd.Short
}
