package voicemail_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("voicemail_users", "SWAGGER_OVERRIDE_/api/v2/voicemail/users")
	voicemail_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_usersCmd)
}

func Cmdvoicemail_users() *cobra.Command {
	return voicemail_usersCmd
}
