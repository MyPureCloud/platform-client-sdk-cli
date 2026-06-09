package telephony_numbers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_numbers_routing"
)

func init() {
	telephony_numbersCmd.AddCommand(telephony_numbers_routing.Cmdtelephony_numbers_routing())
	telephony_numbersCmd.Short = utils.GenerateCustomDescription(telephony_numbersCmd.Short, telephony_numbers_routing.Description, )
	telephony_numbersCmd.Long = telephony_numbersCmd.Short
}
