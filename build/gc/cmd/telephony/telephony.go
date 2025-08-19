package telephony

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony", "SWAGGER_OVERRIDE_/api/v2/telephony")
	telephonyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephonyCmd)
}

func Cmdtelephony() *cobra.Command {
	return telephonyCmd
}
