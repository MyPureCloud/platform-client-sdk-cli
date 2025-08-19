package telephony_providers_edges_phones

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_phones_template"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_phones_reboot"
)

func init() {
	telephony_providers_edges_phonesCmd.AddCommand(telephony_providers_edges_phones_template.Cmdtelephony_providers_edges_phones_template())
	telephony_providers_edges_phonesCmd.AddCommand(telephony_providers_edges_phones_reboot.Cmdtelephony_providers_edges_phones_reboot())
	telephony_providers_edges_phonesCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_phonesCmd.Short, telephony_providers_edges_phones_template.Description, telephony_providers_edges_phones_reboot.Description, )
	telephony_providers_edges_phonesCmd.Long = telephony_providers_edges_phonesCmd.Short
}
