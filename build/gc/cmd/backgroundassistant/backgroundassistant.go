package backgroundassistant

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("backgroundassistant", "SWAGGER_OVERRIDE_/api/v2/backgroundassistant")
	backgroundassistantCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("backgroundassistant"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(backgroundassistantCmd)
}

func Cmdbackgroundassistant() *cobra.Command {
	return backgroundassistantCmd
}
