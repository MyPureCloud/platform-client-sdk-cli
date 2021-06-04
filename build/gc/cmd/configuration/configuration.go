package configuration

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("configuration", "SWAGGER_OVERRIDE_/api/v2/configuration")
	configurationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("configuration"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(configurationCmd)
}

func Cmdconfiguration() *cobra.Command {
	return configurationCmd
}
