package telephony_providers_edges_trunkbasesettings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_trunkbasesettings_template"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_trunkbasesettings_availablemetabases"
)

func init() {
	telephony_providers_edges_trunkbasesettingsCmd.AddCommand(telephony_providers_edges_trunkbasesettings_template.Cmdtelephony_providers_edges_trunkbasesettings_template())
	telephony_providers_edges_trunkbasesettingsCmd.AddCommand(telephony_providers_edges_trunkbasesettings_availablemetabases.Cmdtelephony_providers_edges_trunkbasesettings_availablemetabases())
	telephony_providers_edges_trunkbasesettingsCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_trunkbasesettingsCmd.Short, telephony_providers_edges_trunkbasesettings_template.Description, telephony_providers_edges_trunkbasesettings_availablemetabases.Description, )
	telephony_providers_edges_trunkbasesettingsCmd.Long = telephony_providers_edges_trunkbasesettingsCmd.Short
}