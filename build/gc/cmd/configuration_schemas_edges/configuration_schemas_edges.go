package configuration_schemas_edges

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("configuration_schemas_edges", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas/edges")
	configuration_schemas_edgesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("configuration_schemas_edges"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(configuration_schemas_edgesCmd)
}

func Cmdconfiguration_schemas_edges() *cobra.Command {
	return configuration_schemas_edgesCmd
}
