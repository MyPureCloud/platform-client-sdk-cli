package coaching

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("coaching", "SWAGGER_OVERRIDE_/api/v2/coaching")
	coachingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coachingCmd)
}

func Cmdcoaching() *cobra.Command {
	return coachingCmd
}
