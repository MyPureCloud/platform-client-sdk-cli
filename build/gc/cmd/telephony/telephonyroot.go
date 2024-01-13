package telephony

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_mediaregions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_sipmessages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_siptraces"
)

func init() {
	telephonyCmd.AddCommand(telephony_mediaregions.Cmdtelephony_mediaregions())
	telephonyCmd.AddCommand(telephony_sipmessages.Cmdtelephony_sipmessages())
	telephonyCmd.AddCommand(telephony_providers.Cmdtelephony_providers())
	telephonyCmd.AddCommand(telephony_siptraces.Cmdtelephony_siptraces())
	telephonyCmd.Short = utils.GenerateCustomDescription(telephonyCmd.Short, telephony_mediaregions.Description, telephony_sipmessages.Description, telephony_providers.Description, telephony_siptraces.Description, )
	telephonyCmd.Long = telephonyCmd.Short
}
