package configuration_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("configuration_schemas", "SWAGGER_OVERRIDE_/api/v2/configuration/schemas")
	configuration_schemasCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("configuration_schemas"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(configuration_schemasCmd)
}

func Cmdconfiguration_schemas() *cobra.Command {
	return configuration_schemasCmd
}
