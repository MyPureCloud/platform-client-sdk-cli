package journey_flows_paths

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_flows_paths", "SWAGGER_OVERRIDE_/api/v2/journey/flows/paths")
	journey_flows_pathsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_flows_paths"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_flows_pathsCmd)
}

func Cmdjourney_flows_paths() *cobra.Command {
	return journey_flows_pathsCmd
}
