package telephony_trunkbasesettings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_trunkbasesettings_metrics"
)

func init() {
	telephony_trunkbasesettingsCmd.AddCommand(telephony_trunkbasesettings_metrics.Cmdtelephony_trunkbasesettings_metrics())
	telephony_trunkbasesettingsCmd.Short = utils.GenerateCustomDescription(telephony_trunkbasesettingsCmd.Short, telephony_trunkbasesettings_metrics.Description, )
	telephony_trunkbasesettingsCmd.Long = telephony_trunkbasesettingsCmd.Short
}
