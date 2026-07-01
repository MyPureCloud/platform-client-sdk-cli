package telephony_numbers_routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_numbers_routing_all"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_numbers_routing_reset"
)

func init() {
	telephony_numbers_routingCmd.AddCommand(telephony_numbers_routing_all.Cmdtelephony_numbers_routing_all())
	telephony_numbers_routingCmd.AddCommand(telephony_numbers_routing_reset.Cmdtelephony_numbers_routing_reset())
	telephony_numbers_routingCmd.Short = utils.GenerateCustomDescription(telephony_numbers_routingCmd.Short, telephony_numbers_routing_all.Description, telephony_numbers_routing_reset.Description, )
	telephony_numbers_routingCmd.Long = telephony_numbers_routingCmd.Short
}
