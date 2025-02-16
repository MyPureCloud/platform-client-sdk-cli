package journey_views_data

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_views_data", "SWAGGER_OVERRIDE_/api/v2/journey/views/data")
	journey_views_dataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_views_data"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_views_dataCmd)
}

func Cmdjourney_views_data() *cobra.Command {
	return journey_views_dataCmd
}
