package telephony_providers_edges_autoscalinggroups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_providers_edges_autoscalinggroups", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/autoscalinggroups")
	telephony_providers_edges_autoscalinggroupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_autoscalinggroups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_autoscalinggroupsCmd)
}

func Cmdtelephony_providers_edges_autoscalinggroups() *cobra.Command {
	return telephony_providers_edges_autoscalinggroupsCmd
}
