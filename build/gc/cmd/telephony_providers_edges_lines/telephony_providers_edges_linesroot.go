package telephony_providers_edges_lines

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_lines_template"
)

func init() {
	telephony_providers_edges_linesCmd.AddCommand(telephony_providers_edges_lines_template.Cmdtelephony_providers_edges_lines_template())
	telephony_providers_edges_linesCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_linesCmd.Short, telephony_providers_edges_lines_template.Description, )
	telephony_providers_edges_linesCmd.Long = telephony_providers_edges_linesCmd.Short
}
