package screenmonitors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("screenmonitors", "SWAGGER_OVERRIDE_/api/v2/screenmonitors")
	screenmonitorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("screenmonitors"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(screenmonitorsCmd)
}

func Cmdscreenmonitors() *cobra.Command {
	return screenmonitorsCmd
}
