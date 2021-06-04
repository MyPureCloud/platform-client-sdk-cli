package scripts_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("scripts_uploads", "SWAGGER_OVERRIDE_/api/v2/scripts/uploads")
	scripts_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_uploadsCmd)
}

func Cmdscripts_uploads() *cobra.Command {
	return scripts_uploadsCmd
}
