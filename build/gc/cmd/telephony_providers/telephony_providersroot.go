package telephony_providers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges"
)

func init() {
	telephony_providersCmd.AddCommand(telephony_providers_edges.Cmdtelephony_providers_edges())
	telephony_providersCmd.Short = utils.GenerateCustomDescription(telephony_providersCmd.Short, telephony_providers_edges.Description, )
	telephony_providersCmd.Long = telephony_providersCmd.Short
}
