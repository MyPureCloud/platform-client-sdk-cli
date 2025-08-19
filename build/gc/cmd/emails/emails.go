package emails

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("emails", "SWAGGER_OVERRIDE_/api/v2/emails")
	emailsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("emails"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(emailsCmd)
}

func Cmdemails() *cobra.Command {
	return emailsCmd
}
