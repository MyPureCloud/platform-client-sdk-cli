package telephony_trunkbasesettings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_trunkbasesettings", "SWAGGER_OVERRIDE_/api/v2/telephony/trunkbasesettings")
	telephony_trunkbasesettingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_trunkbasesettings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_trunkbasesettingsCmd)
}

func Cmdtelephony_trunkbasesettings() *cobra.Command {
	return telephony_trunkbasesettingsCmd
}
