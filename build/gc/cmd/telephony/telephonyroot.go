package telephony

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_mediaregions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_sipmessages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_siptraces"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_agents"
)

func init() {
	telephonyCmd.AddCommand(telephony_mediaregions.Cmdtelephony_mediaregions())
	telephonyCmd.AddCommand(telephony_sipmessages.Cmdtelephony_sipmessages())
	telephonyCmd.AddCommand(telephony_providers.Cmdtelephony_providers())
	telephonyCmd.AddCommand(telephony_siptraces.Cmdtelephony_siptraces())
	telephonyCmd.AddCommand(telephony_agents.Cmdtelephony_agents())
	telephonyCmd.Short = utils.GenerateCustomDescription(telephonyCmd.Short, telephony_mediaregions.Description, telephony_sipmessages.Description, telephony_providers.Description, telephony_siptraces.Description, telephony_agents.Description, )
	telephonyCmd.Long = telephonyCmd.Short
}
