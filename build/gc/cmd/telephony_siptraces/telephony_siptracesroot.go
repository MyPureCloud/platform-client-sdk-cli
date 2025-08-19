package telephony_siptraces

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_siptraces_download"
)

func init() {
	telephony_siptracesCmd.AddCommand(telephony_siptraces_download.Cmdtelephony_siptraces_download())
	telephony_siptracesCmd.Short = utils.GenerateCustomDescription(telephony_siptracesCmd.Short, telephony_siptraces_download.Description, )
	telephony_siptracesCmd.Long = telephony_siptracesCmd.Short
}
