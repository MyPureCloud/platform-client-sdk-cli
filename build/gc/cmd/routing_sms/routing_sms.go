package routing_sms

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_sms", "SWAGGER_OVERRIDE_/api/v2/routing/sms")
	routing_smsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_sms"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_smsCmd)
}

func Cmdrouting_sms() *cobra.Command {
	return routing_smsCmd
}
