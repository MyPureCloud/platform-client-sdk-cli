package gdpr

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gdpr", "SWAGGER_OVERRIDE_/api/v2/gdpr")
	gdprCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gdpr"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gdprCmd)
}

func Cmdgdpr() *cobra.Command {
	return gdprCmd
}
