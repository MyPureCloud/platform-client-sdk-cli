package notifications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("notifications", "SWAGGER_OVERRIDE_/api/v2/notifications")
	notificationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("notifications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(notificationsCmd)
}

func Cmdnotifications() *cobra.Command {
	return notificationsCmd
}
