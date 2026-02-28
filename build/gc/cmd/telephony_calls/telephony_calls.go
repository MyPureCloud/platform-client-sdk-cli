package telephony_calls

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_calls", "SWAGGER_OVERRIDE_/api/v2/telephony/calls")
	telephony_callsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_calls"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_callsCmd)
}

func Cmdtelephony_calls() *cobra.Command {
	return telephony_callsCmd
}
