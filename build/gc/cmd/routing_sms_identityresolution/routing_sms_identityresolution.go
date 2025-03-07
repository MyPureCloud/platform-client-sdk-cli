package routing_sms_identityresolution

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_sms_identityresolution", "SWAGGER_OVERRIDE_/api/v2/routing/sms/identityresolution")
	routing_sms_identityresolutionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_sms_identityresolution"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_sms_identityresolutionCmd)
}

func Cmdrouting_sms_identityresolution() *cobra.Command {
	return routing_sms_identityresolutionCmd
}
