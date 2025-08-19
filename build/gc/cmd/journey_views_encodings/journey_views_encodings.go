package journey_views_encodings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_views_encodings", "SWAGGER_OVERRIDE_/api/v2/journey/views/encodings")
	journey_views_encodingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_views_encodings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_views_encodingsCmd)
}

func Cmdjourney_views_encodings() *cobra.Command {
	return journey_views_encodingsCmd
}
