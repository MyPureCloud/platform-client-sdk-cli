package telephony_providers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_providers", "SWAGGER_OVERRIDE_/api/v2/telephony/providers")
	telephony_providersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providersCmd)
}

func Cmdtelephony_providers() *cobra.Command {
	return telephony_providersCmd
}
